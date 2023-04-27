
function sendReply(guildID, channelID, message, client) 
{
    const guild = client.guilds.cache.get(guildID);
    const channel = guild.channels.cache.get(channelID);
    if(channel)
    {
        channel.send(message);
    }

    else
    {
        console.log(`Could not find channel ${channelID} in guild ${guildID} with a name of ${guild.name}`);
    }
}

module.exports = { sendReply };