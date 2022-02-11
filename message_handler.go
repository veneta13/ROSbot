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
				":point_right: use `!log-in ID={CLIENT_ID} SECRET={CLIENT_SECRET}` to connect your Spotify account\n\n" +
				":point_right: use `!say-hi` for a surprise\n\n" +
				":point_right: use `!get-stats {type} {time}` to get your Spotify stats\n\n" +
				":bar_chart: __Supported types__:\n\n" +
				":singer: artist\n\n" +
				":musical_note: track\n\n" +
				":alarm_clock: __Supported time periods__:\n\n" +
				":clock1: last month\n\n" +
				":clock2: 6 months\n\n" +
				":clock3: all time\n\n" +
				":point_right: use `!create-playlist {mood}` to create a playlist\n\n" +
				":performing_arts: __Supported moods__:\n\n" +
				":smile: happy\n\n" +
				":sob: sad\n\n" +
				":relaxed: relaxed\n\n" +
				":partying_face: party\n\n" +
				":face_with_monocle: focused\n\n" +
				":smiling_face_with_3_hearts: romantic\n\n" +
				":christmas_tree: holiday\n\n" +
				":blue_car: travel\n\n" +
				":raised_hands: motivated \n\n" +
				":sleeping: sleepy \n\n",
			},
		}
		_, _ = session.ChannelMessageSendComplex( message.ChannelID, messageContent)
	}

	if strings.Contains(message.Content, "!log-in"){

		if strings.Contains(message.Content, "ID=") && strings.Contains(message.Content, "SECRET="){
			clientID = strings.SplitAfter(strings.Split(message.Content, " ")[1], "ID=")[1]
			clientSecret = strings.SplitAfter(message.Content, "SECRET=")[1]
			deleteMessage(session, message)
		} else {
			fmt.Println("Cannot log in")
			_, _ = session.ChannelMessage( message.ChannelID, "Unsuccessful Spotify login :cry:")
			deleteMessage(session, message)
			return
		}
		auth = New(
			WithClientID(clientID),
			WithClientSecret(clientSecret),
			WithRedirectURL(redirectUri),
			WithScopes(
				ScopeUserReadPrivate,
				ScopePlaylistModifyPublic,
				ScopePlaylistModifyPrivate,
				ScopeUserLibraryModify,
				ScopeUserLibraryRead,
				ScopeUserTopRead,
				ScopeUserModifyPlaybackState,
				ScopeImageUpload))
		loginLink := spotifyLogin()
		_, _ = session.ChannelMessageSend(message.ChannelID, loginLink)

		client = <-ch
		privateUser, _ := client.CurrentUser(context.Background())
		user = privateUser

		fmt.Println("Log: Logged in as:", user.ID)
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

			playlist, err = getPlaylistByMood(message.Content)

			if err != nil {
				_, _ = session.ChannelMessageSend(
					message.ChannelID,
					"Cannot create playlist :pensive: Please try again")
				fmt.Println(err)
				return
			}

			if playlist == nil {
				_, _ = session.ChannelMessageSend(
					message.ChannelID,
					"Key word not recognised :cry: Please try again")
				return
			}

			fmt.Println("Log: " + playlist.Name + " created successfully ")

			messageContent := &discordgo.MessageSend{
				Content: playlist.Name + " created successfully :partying_face:",
				Embed: &discordgo.MessageEmbed{
					Image: &discordgo.MessageEmbedImage{
						URL: playlistCoverURL,
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

	// !get-stats command
	if strings.Contains(message.Content, "!get-stats"){
		if user != nil {
			types, time := getStatsType(message.Content)
			tracks, artists, err := getStats(types, time)

			if err != nil {
				fmt.Println("Error reading user stats")
				_, _ = session.ChannelMessageSend(message.ChannelID, "Getting your stats was unsuccessful :cry: Please try again")
				return
			}

			if tracks == nil && artists == nil {
				fmt.Println("Unrecognised command")
				_, _ = session.ChannelMessageSend(message.ChannelID, "I don't recognise this command :thinking: Please try again")
				return
			}

			if tracks != nil {
				trackList := makeTrackList(tracks)

				messageContent := &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Image:  &discordgo.MessageEmbedImage{
							URL: trackList[0].Album.Images[0].URL,
						},
						Color: 0xffd700,
						Description:
						":trophy: **YOUR TOP SONGS**\n\n" +
						":first_place: " + trackList[0].Name + " - " + trackList[0].Artists[0].Name + "\n" +
						":second_place: " + trackList[1].Name + " - " + trackList[1].Artists[0].Name + "\n" +
						":third_place: " + trackList[2].Name + " - " + trackList[2].Artists[0].Name + "\n" +
						"4. " + trackList[3].Name + " - " + trackList[3].Artists[0].Name + "\n" +
						"5. " + trackList[4].Name + " - " + trackList[4].Artists[0].Name + "\n" +
						"6. " + trackList[5].Name + " - " + trackList[5].Artists[0].Name + "\n" +
						"7. " + trackList[6].Name + " - " + trackList[6].Artists[0].Name + "\n" +
						"8. " + trackList[7].Name + " - " + trackList[7].Artists[0].Name + "\n" +
						"9. " + trackList[8].Name + " - " + trackList[8].Artists[0].Name + "\n" +
						"10. " + trackList[9].Name + " - " + trackList[9].Artists[0].Name + "\n",
					},
				}
				_, _ = session.ChannelMessageSendComplex( message.ChannelID, messageContent)
			}

			if artists != nil {
				artistList := makeArtistList(artists)

				messageContent := &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Image:  &discordgo.MessageEmbedImage{
							URL: artistList[0].Images[0].URL,
						},
						Color: 0xffd700,
						Description:
						":trophy: **YOUR TOP ARTISTS**\n\n" +
							":first_place: " + artistList[0].Name + "\n" +
							":second_place: " + artistList[1].Name + "\n" +
							":third_place: " + artistList[2].Name + "\n" +
							"4. " + artistList[3].Name + "\n" +
							"5. " + artistList[4].Name + "\n",
					},
				}
				_, _ = session.ChannelMessageSendComplex( message.ChannelID, messageContent)
			}
			return

		} else {
			fmt.Println("Log: Require login")
			_, _ = session.ChannelMessageSend(message.ChannelID, "Please `!log-in` go get your stats :wink:")
		}
	}
}

func deleteMessage (session *discordgo.Session, message *discordgo.MessageCreate) {
	err := session.ChannelMessageDelete(message.ChannelID, message.ID)

	if err != nil {
		fmt.Println("Cannot delete message")
		_, _ = session.ChannelMessageSend(
			message.ChannelID,
			"<@!" + message.Author.ID + "> Please delete your login message:exclamation:")
	}
}
