# TeleCord

This software creates a two-way interface between Telegram and Discord.

## How to use
**The tunnel must be opened on BOTH receiving and sending end !**
### Receiving Messages
#### On Discord
Join [this discord server](https://discord.gg/dmMev8d) and open your Direct Messages / Send a DM to TeleCord.
#### On Telegram
Start a conversation with [this bot](https://t.me/millefeuilleTeleCordBot).

### Sending Messages
#### From Discord
Send `/dest -[Contact telegramID]` to the TeleCord bot and the tunnel `Discord->Telegram` is now open.\
To get your userID you need to have dev mode activated in Discord and then right click the desired profile, and select `Copy ID`
#### From Telegram
Send `/dest -[Contact discordID]` to the TeleCord bot and the tunnel `Telegram->Discord` is now open.\
To get your userID start a convo with `userinfobot` on Telegram.

## Known Bugs
- Attachments such as images, audio, files... Are not working

## Roadmap
- Adding the userID next to the name
- Adding a command to get your own userID
- Adding a command to know current destination
- Adding attachments support
- Adding more platforms (Slack for example)
