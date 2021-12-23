package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func startServer() {
	// Create Discord session
	var sessionError error
	discordSession, sessionError = discordgo.New("Bot " + DiscordToken)
	if sessionError != nil {
		fmt.Println("Error: Cannot create Discord session,", sessionError)
		return
	}

	// messageCreate - callback for MessageCreate events
	discordSession.AddHandler(messageCreate)

	// receiving message events
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	// open websocket connection
	sessionError = discordSession.Open()
	if sessionError != nil {
		fmt.Println("Error: cannot open connection,", sessionError)
		return
	}

	fmt.Println("Log: ROSbot is up.")
}

func stopServer(){
	err := discordSession.Close()
	if err != nil {
		fmt.Println("Error: error closing session,", err)
		return
	}
}
