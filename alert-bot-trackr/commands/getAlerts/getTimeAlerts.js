const { SlashCommandBuilder } = require('discord.js');
const { TimeAlerts } = require ('../../db/database.js');


module.exports = {
	data: new SlashCommandBuilder()
		.setName('get-time-alerts')
		.setDescription('Outputs all the connection alerts.'),
            
	async execute(interaction) {
        const timeAlerts = await TimeAlerts.findAll();
        if(timeAlerts.length != 0)
        {
            output = '';
            for (const timeAlert of timeAlerts) {
                output += `Alert ID: \`${timeAlert.id}\` | Username: \`${timeAlert.username}\` | Time: \`${timeAlert.setTime}\` | Message: \`${timeAlert.message}\`\n`;
            }
            return interaction.reply(output);
        }

        return interaction.reply(`There are no time alerts.`);
	},
};