const {TimeAlerts, ThresholdAlerts, ConnectionAlerts} = require ('./db/database.js');
const axios = require('axios');
const { API_ENDPOINT, limit, order } = require('./config.json');
const { sendReply } = require('./utils/sendReply.js');
const {deleteEntry} = require('./utils/deleteEntry.js');


async function alertLogic(){
    const client = require('./index.js'); //if outside of the function, client is undefined
    try
    {
        //first process time alerts
        const timeAlerts = await TimeAlerts.findAll();
        if(timeAlerts.length != 0)
        {
            for(const timeAlert of timeAlerts)
            {
                const id = timeAlert.id;
                const username = timeAlert.username;
                const message = timeAlert.message;
                const setTime = timeAlert.setTime;
                const timeStamp = timeAlert.timeStamp;
                const channelID = timeAlert.channelID;
                const guildID = timeAlert.guildID;
                const userID = timeAlert.userID;

                if( (Math.floor(Date.now() / 1000) - timeStamp) >= setTime )
                {
                    var taggedMessage = `<@${userID}> **TIME ALERT:** --- ${message}`;
                    sendReply(guildID, channelID, taggedMessage, client);
                    await TimeAlerts.destroy({ where: { id } });
                }
            }
        }

        //get all the threshold alerts
        const thresholdAlerts = await ThresholdAlerts.findAll();
        //get all the connection alerts
        const connectionAlerts = await ConnectionAlerts.findAll();

        if(thresholdAlerts.length != 0) //TODO make it only run for new values
        {
            for(const thresholdAlert of thresholdAlerts)
            {
                const fieldID = thresholdAlert.fieldID;
                const apiKey = thresholdAlert.apiKey;
                const min = thresholdAlert.thresholdMin;
                const max = thresholdAlert.thresholdMax;
                const offset = thresholdAlert.totalValues; //make offset the total values saved in the database, so that we only take new values
                const channelID = thresholdAlert.channelID;
                const guildID = thresholdAlert.guildID;

                // console.log(`checking threshold alert with api key "${apiKey}" and field id "${fieldID}" for new values outside the threshold of ${min} or/and ${max}.`)

                const response = await axios.get(`${API_ENDPOINT}?apiKey=${apiKey}&fieldId=${fieldID}&offset=${offset}&limit=${limit}&order=${order}`).
                catch(error => 
                    {
                        if(error.response && error.response.status == 400)
                        {
                            console.log(error.response.data.error);
                            if(error.response.data.error == "Failed to find field." || error.response.data.error == "Failed to find project, invalid API key.")
                            {
                                //delete entry in all the fields with a combination of that api and field id
                                console.log(`deleting entry with api key "${apiKey}" and field id "${fieldID}" since it does not exist.`);
                                deleteEntry(apiKey, fieldID);
                            }
                            return
                        }

                        else //TODO figure out what to do when server is down or error message 500
                        {
                            //send error message
                            console.log(error);
                        }
                    });
                
                if(response)
                {

                    const values = response.data.values;
                    const totalValues = response.data.totalValues;

                    //offset 0 means that this is the first time the alert is being run
                    if(values.length != 0 && offset != 0)
                    {
                        var responseValues = [];
                        
                        for(const entry of values)
                        {
                            var value = entry.value;
                            if(min === null && max === null)
                            {
                                responseValues.push(value);
                            }
                            
                            else if( (min !== null && value < min) || (max !== null && value > max) )
                            {
                                responseValues.push(value);
                            }
                        }

                        //send the message only if values are outside the thresholds
                        if(responseValues.length !== 0)
                        {
                            var message = `@everyone THRESHOLD ALERT: New values [\`${responseValues}\`] received that is/are outside the threshold of ${min} or/and ${max}.\n\`API Key: ${apiKey}\`\n\`Field ID: ${fieldID}\``;
                            sendReply(guildID, channelID, message, client)
                        }
                    }

                    //update the totalValues, so that offset is correct
                    await ThresholdAlerts.update({ totalValues: totalValues }, { where: { apiKey: apiKey, fieldID:fieldID } });
                }
            }
        }

        if(connectionAlerts.length != 0)
        {
            for(const connectionAlert of connectionAlerts)
            {
                const fieldID = connectionAlert.fieldID;
                const apiKey = connectionAlert.apiKey;
                const offset = connectionAlert.totalValues; //make offset the total values saved in the database, so that we only take new values
                const channelID = connectionAlert.channelID;
                const guildID = connectionAlert.guildID;
                const timeStamp = connectionAlert.timeStamp;
                const setTime = connectionAlert.setTime;

                // console.log(`Checking connection alert with api key "${apiKey}" and field id "${fieldID}"`)

                const response = await axios.get(`${API_ENDPOINT}?apiKey=${apiKey}&fieldId=${fieldID}&offset=${offset}&limit=${limit}&order=${order}`).
                catch(error => 
                    {
                        if(error.response && error.response.status == 400)
                        {
                            console.log(error.response.data.error);
                            if(error.response.data.error == "Failed to find field." || error.response.data.error == "Failed to find project, invalid API key.")
                            {
                                //delete entry in all the fields with a combination of that api and field id
                                console.log(`deleting entry with api key "${apiKey}" and field id "${fieldID}" since it does not exist.`);
                                deleteEntry(apiKey, fieldID);
                            }
                            return
                        }

                        else //TODO figure out what to do when server is down or error message 500
                        {
                            //send error message
                            console.log(error);
                        }
                    });
                
                if(response)
                {

                    const values = response.data.values;
                    const totalValues = response.data.totalValues;

                    //offset 0 means that this is the first time the alert is being run
                    if(values.length != 0 && offset != 0)
                    {
                        //we got some values so need to update the timestamp
                        await ConnectionAlerts.update({ timeStamp: Math.floor(Date.now() / 1000) }, { where: { apiKey: apiKey, fieldID:fieldID } });
                    }

                    else if(values.length == 0)
                    {
                        //no values so check if the time has passed, if the time has passed then might be a connection issue
                        if( (Math.floor(Date.now() / 1000) - timeStamp) >= setTime )
                        {
                            var message = `@everyone CONNECTION ALERT: No values received for the last \`${setTime}\` seconds.\n\`API Key: ${apiKey}\`\n\`Field ID: ${fieldID}\``;
                            sendReply(guildID, channelID, message, client)

                            //reset the timestamp so that alert is issued again in case the issue is not solved
                            await ConnectionAlerts.update({ timeStamp: Math.floor(Date.now() / 1000) }, { where: { apiKey: apiKey, fieldID:fieldID } });
                        }
                    }

                    //update the totalValues, so that offset is correct
                    await ConnectionAlerts.update({ totalValues: totalValues }, { where: { apiKey: apiKey, fieldID:fieldID } });
                }
            }
        }

    }
    catch(error)
    {
        console.log(error);
    }
};
module.exports = {alertLogic};