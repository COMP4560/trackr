const { SlashCommandBuilder } = require('discord.js');
const { ConnectionAlerts } = require ('../../db/database.js');


module.exports = {
	data: new SlashCommandBuilder()
		.setName('get-connection-alerts')
		.setDescription('Outputs all the connection alerts.'),
            
	async execute(interaction) {
        const connectionAlerts = await ConnectionAlerts.findAll();
        if(connectionAlerts.length != 0)
        {
            output = '';
            for (const connectionAlert of connectionAlerts) {
                output += `Alert ID: \`${connectionAlert.id}\` | APIkey: \`${connectionAlert.apiKey}\` | FieldID: \`${connectionAlert.fieldID}\` | Username: \`${connectionAlert.username}\` | Time: \`${connectionAlert.setTime}\`\n`;
            }
            return interaction.reply(output);
        }

        return interaction.reply(`There are no connection alerts.`);
	},
};