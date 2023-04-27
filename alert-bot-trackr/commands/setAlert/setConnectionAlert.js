// a time alert, that acts basically as an alarm or a reminder. A user
// can set a time alert with a specific timer and a message, and when the timer is up,
// the bot will echo a message to the user. After reminding, the aler is deleted from DB.


const { SlashCommandBuilder } = require('discord.js');
const { ConnectionAlerts, APIs } = require ('../../db/database.js');


module.exports = {
	data: new SlashCommandBuilder()
		.setName('set-connection-alert')
		.setDescription('Set a connection alert with a timeframe in which you expect to receive values.')
        .addStringOption(option =>
            option.setName('api')
                .setDescription('your project API key')
                .setRequired(true))
        .addIntegerOption(option =>
            option.setName('field-id')
                .setDescription('FieldID you want alert for')
                .setRequired(true))
        .addNumberOption(option =>
            option.setName('time')
            .setDescription('Time in seconds that the reminder will be set for')
            .setRequired(true)),
            
	async execute(interaction) {
		const apiKey = interaction.options.getString('api');
        const fieldID = interaction.options.getInteger('field-id');
        const time = interaction.options.getNumber('time');
        const guildID = String(interaction.guild.id);
        const channelID = String(interaction.channel.id);
        const username = interaction.user.username;
        const userID = interaction.user.id;

        const api = await APIs.findOrCreate({
            where: {api:apiKey},
            defaults: { api:apiKey}
        });

        try{
            const connectionAlert = await ConnectionAlerts.create({
                apiKey: apiKey,
                fieldID: fieldID,
                setTime: time,
                timeStamp: Math.floor(Date.now() / 1000), // current time in seconds
                guildID: guildID,
                channelID: channelID,
                username: username,
                userID: userID,
            });

            return interaction.reply(`Connection alert with an id of \`${connectionAlert.id}\` has been set for \`${time}\` seconds.`);
        }

        catch(error)
        {
            if (error.name === 'SequelizeUniqueConstraintError') {
				return interaction.reply(`**Alert with API key: \`${apiKey}\` and field ID: \`${fieldID}\` already exists. Delete previous alert and create a new one.**`);
			}
            else{
                return interaction.reply('Something went wrong with adding a connection alert.');
            }
        }


		
	},
};