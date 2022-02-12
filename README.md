
# ROSbot

> A Discord bot generating Spotify playlists

## Setup

First, you need to get a Discord bot token from [here](https://discord.com/developers/applications).  
You need to get a Client ID and Client Secret tokens from [here](https://developer.spotify.com/dashboard/).  
  
Next, make sure that nothing is running on <http://localhost:8888/>  
To run the bot use the following command:

```shell
go run . -t DISCORD_BOT_TOKEN
```

The following message should appear on the command prompt:

```shell
Log: rosbot is up.
```

## Commands

ROSbot supports the following commands:

| Command | Description | Arguments |
|:-:|:-:|:-:|
| `!help` | print a list of all commands |   |
| `!log-in  ID={client_id} SECRET={client_secret}` | print a link for Spotify authorisation | `client_id` - the Client ID from Spotify `client_secret` - the Client Secret from Spotify |
| `!say-hi` | greet the user |   |
| `!get-stats {type} {time}` | get the Spotify statistics of the user | `{type}` - artist / track `{time}` - last month / last 6 months / all time |
|`!create-playlist {mood}`| create a Spotify playlist according to a given mood | `{mood}` - happy / sad / relaxed / party / focused / romantic / holiday / travel / motivated / sleepy |
