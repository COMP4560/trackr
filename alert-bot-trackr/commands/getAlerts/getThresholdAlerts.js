const { SlashCommandBuilder } = require('discord.js');
const { ThresholdAlerts } = require ('../../db/database.js');


module.exports = {
	data: new SlashCommandBuilder()
		.setName('get-threshold-alerts')
		.setDescription('Outputs all the treshold alerts.'),
            
	async execute(interaction) {
        const thresholdAlerts = await ThresholdAlerts.findAll();
        if(thresholdAlerts.length != 0)
        {
            output = '';
            for (const thresholdAlert of thresholdAlerts) {
                output += `Alert ID: \`${thresholdAlert.id}\` | APIkey: \`${thresholdAlert.apiKey}\` | FieldID: \`${thresholdAlert.fieldID}\` | Username: \`${thresholdAlert.username}\` | Min: \`${thresholdAlert.thresholdMin}\` | Max: \`${thresholdAlert.thresholdMax}\`\n`;
            }
            return interaction.reply(output);
        }

        return interaction.reply(`There are no threshold alerts.`);
	},
};