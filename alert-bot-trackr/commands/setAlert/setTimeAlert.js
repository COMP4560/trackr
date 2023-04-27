// a time alert, that acts basically as an alarm or a reminder. A user
// can set a time alert with a specific timer and a message, and when the timer is up,
// the bot will echo a message to the user. After reminding, the aler is deleted from DB.


const { SlashCommandBuilder } = require('discord.js');
const { TimeAlerts } = require ('../../db/database.js');



module.exports = {
	data: new SlashCommandBuilder()
		.setName('set-time-alert')
		.setDescription('Set a reminder with a message to be sent after a specific time. Reminder is set only once.')
        .addStringOption(option =>
            option.setName('message')
                .setDescription('Reminder message')
                .setRequired(true))
        .addNumberOption(option =>
            option.setName('time')
            .setDescription('Time in seconds that the reminder will be set for')
            .setRequired(true)),
            
	async execute(interaction) {
		const message = interaction.options.getString('message');
        const time = interaction.options.getNumber('time');
        const guildID = interaction.guild.id;
        const channelID = interaction.channel.id;
        const username = interaction.user.username;
        const userID = interaction.user.id;

        try{
            const timeAlert = await TimeAlerts.create({
                username: interaction.user.username,
                message: message,
                setTime: time,
                timeStamp: Math.floor(Date.now() / 1000), // current time in seconds
                guildID: guildID,
                channelID: channelID,
                username: username,
                userID: userID,
            });

            return interaction.reply(`time alert with an id of \`${timeAlert.id}\` has been set for \`${time}\` seconds with a message: \`${message}\`.`);
        }

        catch(error)
        {
            return interaction.reply('Something went wrong with adding a time alert.');
        }


		
	},
};