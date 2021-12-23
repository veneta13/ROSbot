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
	if message.Content == "!say-hi" && user != nil{
		_, _ = session.ChannelMessageSend(message.ChannelID, "Hello")
	}

	// !create-playlist command
	if strings.Contains(message.Content, "!create-playlist") && user != nil {
		var err error
		var playlist *spotify.FullPlaylist

		if strings.Contains(message.Content, "happy"){
			playlist, err = createPlaylist("happy")
		}

		if strings.Contains(message.Content, "sad"){
			playlist, err = createPlaylist("sad")
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Log: " + playlist.Name + " created successfully")
	}
}
