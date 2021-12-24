package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
	"strings"
)

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!help" {
		messageContent := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Color: 0x0088de,
				Description:
				":scroll: **COMMANDS**\n\n" +
					":point_right: use `!help` to get this list\n\n" +
					":point_right: use `!log-in` to connect your Spotify account\n\n" +
					":point_right: use `!create-playlist {mood}` to create a playlist\n\n" +
					":grey_exclamation: **Supported moods** :\n\n" +
					":smile: happy\n\n" +
					":point_right: use `!say-hi` for a surprise\n",
			},
		}
		_, _ = session.ChannelMessageSendComplex( message.ChannelID, messageContent)
	}

	if message.Content == "!log-in"{
		loginLink := spotifyLogin()
		_, _ = session.ChannelMessageSend(message.ChannelID, loginLink)

		client = <-ch
		privateUser, _ := client.CurrentUser(context.Background())
		user = privateUser

		fmt.Println("Log: Logged in as:", user.ID)
		return
	}

	// !say-hi command
	if message.Content == "!say-hi"{
		_, _ = session.ChannelMessageSend(message.ChannelID, "Hello " + message.Author.Username + " :wave:")
	}

	// !create-playlist command
	if strings.Contains(message.Content, "!create-playlist"){
		if user != nil {
			var playlist *spotify.FullPlaylist
			var err error

			if strings.Contains(message.Content, "happy"){
				playlist, err = makeCompletePlaylist("happy", PlaylistCoverFile)
				if err != nil {
					return
				}
			}

			if err != nil {
				_, _ = session.ChannelMessageSend(
					message.ChannelID,
					"Cannot create playlist :pensive: Please try again")
				fmt.Println(err)
			}

			fmt.Println("Log: " + playlist.Name + " created successfully ")

			messageContent := &discordgo.MessageSend{
				Content: "Playlist created successfully :partying_face:",
				Embed: &discordgo.MessageEmbed{
					Image: &discordgo.MessageEmbedImage{
						URL: PlaylistCoverURL,
					},
					Color: 0x0088de,
					Description:
					"You can access your playlist here :point_right:" +
						"https://open.spotify.com/playlist/" +
						string(playlist.ID),
				},
			}
			_, _ = session.ChannelMessageSendComplex( message.ChannelID, messageContent)

		} else {
			fmt.Println("Log: Require login")
			_, _ = session.ChannelMessageSend(message.ChannelID, "Please `!log-in` before creating playlists :wink:")

		}
	}
}
