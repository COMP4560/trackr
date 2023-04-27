// a time alert, that acts basically as an alarm or a reminder. A user
// can set a time alert with a specific timer and a message, and when the timer is up,
// the bot will echo a message to the user. After reminding, the aler is deleted from DB.


const { SlashCommandBuilder } = require('discord.js');
const { ThresholdAlerts } = require ('../../db/database.js');



module.exports = {
	data: new SlashCommandBuilder()
		.setName('delete-threshold-alert')
		.setDescription('Delete a threshold alert with a specific id.')
        .addIntegerOption(option =>
            option.setName('id')
            .setDescription('threshold alert id')
            .setRequired(true)),
            
	async execute(interaction) {
        const id = interaction.options.getInteger('id');
        const rowCount = await ThresholdAlerts.destroy({ where: { id: id } });
        if (!rowCount) return interaction.reply(`Threshold alert with an id of \`${id}\` does not exist.`);

        return interaction.reply(`Threshold alert with an id of \`${id}\` has been deleted.`);
	},
};