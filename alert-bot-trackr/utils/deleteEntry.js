//deletes entry with given API key value and field ID value from all tables
const {ThresholdAlerts, ConnectionAlerts} = require ('../db/database.js');
async function deleteEntry(apiKey, fieldID)
{
    // Delete records from the `users` table
    await ThresholdAlerts.destroy({ where: { apiKey, fieldID } });

    // Delete records from the `posts` table
    await ConnectionAlerts.destroy({ where: { apiKey, fieldID } });
}

module.exports = {deleteEntry};