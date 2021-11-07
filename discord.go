package main

import (
	"github.com/bwmarrin/discordgo"
)

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	// !say-hi command
	if message.Content == "!say-hi" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Hello ")
	}

	// !create-playlist command
	if message.Content == "!create-playlist" {

	}
}
