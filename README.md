
# ROSbot

A Discord bot generating Spotify playlists

## Setup

First, you need to get a Discord bot token from [here](https://discord.com/developers/applications).  
You need to get a Client ID and Client Secret tokens from [here](https://developer.spotify.com/dashboard/).  
  
Next, make sure that nothing is running on <http://localhost:PORT/>.
To run the bot use the following command:

```shell
go run . -t DISCORD_BOT_TOKEN -port PORT
```

You can omit the port, as the default is 8888.

The following message should appear on the command prompt if the bot is set up correctly:

```shell
Log: ROSbot is up.
```

## Commands

ROSbot supports the following commands:

| Command | Description | Arguments | Image |
|:-:|:-:|:-:|:-:|
| `!help` | print a list of all commands |   |![Help command](/readme_pics/command.PNG) |
| `!log-in  ID={client_id} SECRET={client_secret}` | print a link for Spotify authorisation | `client_id` - the Client ID from Spotify `client_secret` - the Client Secret from Spotify |   |
| `!say-hi` | greet the user |   |   |
| `!get-stats {type} {time}` | get the Spotify statistics of the user | `{type}` - artist / track `{time}` - last month / last 6 months / all time | ![Artist statistics](/readme_pics/artists.png) ![Track statistics](/readme_pics/songs.png) |
|`!create-playlist {mood}`| create a Spotify playlist according to a given mood | `{mood}` - happy / sad / relaxed / party / focused / romantic / holiday / travel / motivated / sleepy | ![Playlist](/readme_pics/playlist.png) |

