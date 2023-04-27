const { Events } = require('discord.js');
const { TimeAlerts, ThresholdAlerts, ConnectionAlerts, APIs} = require ('../db/database.js');

module.exports = {
	name: Events.ClientReady,
	once: true,
	execute(client) {
		//sync the database
		APIs.sync();
		TimeAlerts.sync();
		ThresholdAlerts.sync();
		ConnectionAlerts.sync();
		console.log(`Ready! Logged in as ${client.user.tag}`);
	},
};