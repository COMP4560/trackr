 # **Trackr Alert Bot (TAB)**
 - Alert system for the Trackr (https://github.com/University-of-Manitoba-Computer-Science/trackr)
 - Has 3 types of alers (Threshold, Connection, and Time)
   - Threshold: Alerts when a field value is above or below a certain value
   - Connection: Alerts when a device has not sent data in a certain amount of time
   - Time: works just like a reminder, alerts after a certain amount of time has passed once and sends a user-specified message

## **To use**:
- Deploy commands (need to do only once if commands were not changed):
  - `node .\deploy-commands.js`
- Run the bot:
  - `node .` or `node .\index.js`



## **Things to do in the future**:
- add validation for creation of alerts (checking api and fields exist) -> for now it validates api keys and field ids only when processing alerts and fetching data

- get rid of API table in the database -> was made with some design decisions in mind that are no longer relevant and therefor is not needed

- add more alert types (e.g. if a field value is equal to a certain value)

- add more ways to send alerts (e.g. email, text, etc.)

- **Make the Trackr send updates to the bot for the projects that have alerts set up. For now, bot gets the values from the projects that have alerts set up every second and then processes alerts -> creates a lot of get requests per second if there are a lot of alerts (therefore I limited it to one alert of each kind per project)**
 