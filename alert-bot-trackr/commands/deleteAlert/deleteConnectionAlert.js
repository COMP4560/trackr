// a time alert, that acts basically as an alarm or a reminder. A user
// can set a time alert with a specific timer and a message, and when the timer is up,
// the bot will echo a message to the user. After reminding, the aler is deleted from DB.


const { SlashCommandBuilder } = require('discord.js');
const { ConnectionAlerts } = require ('../../db/database.js');



module.exports = {
	data: new SlashCommandBuilder()
		.setName('delete-connection-alert')
		.setDescription('Delete a connection alert with a specific id.')
        .addIntegerOption(option =>
            option.setName('id')
            .setDescription('threshold alert id')
            .setRequired(true)),
            
	async execute(interaction) {
        const id = interaction.options.getInteger('id');
        const rowCount = await ConnectionAlerts.destroy({ where: { id: id } });
        if (!rowCount) return interaction.reply(`Connection alert with an id of \`${id}\` does not exist.`);

        return interaction.reply(`Connection alert with an id of \`${id}\` has been deleted.`);
	},
};