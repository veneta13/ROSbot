package rosbot

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
	"strings"
)

// Handles input messages and creates a response.
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!help" {
		messageContent := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Color:       0x0088de,
				Description: logger(1),
			},
		}
		_, _ = session.ChannelMessageSendComplex(message.ChannelID, messageContent)
	}

	if strings.Contains(message.Content, "!log-in") {
		var clientID string

		var clientSecret string

		if strings.Contains(message.Content, "ID=") && strings.Contains(message.Content, "SECRET=") {
			clientID, clientSecret = getClientCredentials(message.Content)
			deleteMessage(session, message)
		} else {
			commandLineLogger(6)
			_, _ = session.ChannelMessage(message.ChannelID, "Unsuccessful Spotify login :cry:")
			deleteMessage(session, message)

			return
		}

		auth = New(
			WithClientID(clientID),
			WithClientSecret(clientSecret),
			WithRedirectURL(projectProperties.redirectURL),
			WithScopes(
				ScopeUserReadPrivate,
				ScopePlaylistModifyPublic,
				ScopePlaylistModifyPrivate,
				ScopeUserLibraryModify,
				ScopeUserLibraryRead,
				ScopeUserTopRead,
				ScopeUserModifyPlaybackState,
				ScopeImageUpload))
		loginLink := SpotifyLogin()
		_, _ = session.ChannelMessageSend(message.ChannelID, loginLink)

		users[message.Author.ID] = <-ch

		commandLineLogger(7)
	}

	// !say-hi command
	if message.Content == "!say-hi" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Hello "+message.Author.Username+" :wave:")
	}

	// !create-playlist command
	if strings.Contains(message.Content, "!create-playlist") {
		client := getClient(message.Author.ID)
		user, _ := client.CurrentUser(context.Background())

		if user != nil {
			var playlist *spotify.FullPlaylist

			var err error

			playlist, err = GetPlaylistByMood(message.Content, client)

			if err != nil {
				_, _ = session.ChannelMessageSend(
					message.ChannelID,
					"Cannot create playlist :pensive: Please try again")

				commandLineLogger(19)

				return
			}

			if playlist == nil {
				_, _ = session.ChannelMessageSend(
					message.ChannelID,
					"Key word not recognised :cry: Please try again")

				return
			}

			commandLineLogger(8)

			messageContent := &discordgo.MessageSend{
				Content: playlist.Name + " created successfully :partying_face:",
				Embed: &discordgo.MessageEmbed{
					Image: &discordgo.MessageEmbedImage{
						URL: projectProperties.playlistCoverURL,
					},
					Color: 0x0088de,
					Description: "You can access your playlist here :point_right:" +
						"https://open.spotify.com/playlist/" +
						string(playlist.ID),
				},
			}
			_, _ = session.ChannelMessageSendComplex(message.ChannelID, messageContent)
		} else {
			fmt.Println("Log: Require login")
			_, _ = session.ChannelMessageSend(message.ChannelID, "Please `!log-in` before creating playlists :wink:")
		}
	}

	// !get-stats command
	if strings.Contains(message.Content, "!get-stats") {
		client := getClient(message.Author.ID)
		user, _ := client.CurrentUser(context.Background())

		if user != nil {
			types, time := getStatsType(message.Content)
			tracks, artists, err := getStats(types, time, client)

			if err != nil {
				commandLineLogger(9)

				_, _ = session.ChannelMessageSend(message.ChannelID, "Getting your stats was unsuccessful :cry: Please try again")

				return
			}

			if tracks == nil && artists == nil {
				commandLineLogger(10)

				_, _ = session.ChannelMessageSend(message.ChannelID, "I don't recognise this command :thinking: Please try again")

				return
			}

			if tracks != nil {
				trackList := makeTrackList(tracks)

				messageContent := &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Image: &discordgo.MessageEmbedImage{
							URL: trackList[0].Album.Images[0].URL,
						},
						Color: 0xffd700,
						Description: ":trophy: **YOUR TOP SONGS**\n\n" +
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
				_, _ = session.ChannelMessageSendComplex(message.ChannelID, messageContent)
			}

			if artists != nil {
				artistList := makeArtistList(artists)

				messageContent := &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Image: &discordgo.MessageEmbedImage{
							URL: artistList[0].Images[0].URL,
						},
						Color: 0xffd700,
						Description: ":trophy: **YOUR TOP ARTISTS**\n\n" +
							":first_place: " + artistList[0].Name + "\n" +
							":second_place: " + artistList[1].Name + "\n" +
							":third_place: " + artistList[2].Name + "\n" +
							"4. " + artistList[3].Name + "\n" +
							"5. " + artistList[4].Name + "\n",
					},
				}
				_, _ = session.ChannelMessageSendComplex(message.ChannelID, messageContent)
			}

			return
		}

		commandLineLogger(11)

		_, _ = session.ChannelMessageSend(message.ChannelID, "Please `!log-in` go get your stats :wink:")
	}
}

// Delete the user login message after retrieving the login credentials.
func deleteMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	err := session.ChannelMessageDelete(message.ChannelID, message.ID)

	if err != nil {
		commandLineLogger(12)

		_, _ = session.ChannelMessageSend(
			message.ChannelID,
			"<@!"+message.Author.ID+"> Please delete your login message:exclamation:")
	}
}

// Get Spotify client ID and secret key from the user login message.
func getClientCredentials(message string) (clientID string, clientSecret string) {
	clientID = strings.SplitAfter(strings.Split(message, " ")[1], "ID=")[1]
	clientSecret = strings.SplitAfter(message, "SECRET=")[1]

	return
}

// Get Spotify client from Discord user ID.
func getClient(discordID string) *spotify.Client {
	if val, ok := users[discordID]; ok {
		return val
	}

	commandLineLogger(13)

	return nil
}
