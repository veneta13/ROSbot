package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	DiscordToken string
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "ROSbot Token")
	flag.Parse()
}

func main() {

	// Create Discord session
	discordSession, sessionError := discordgo.New("Bot " + DiscordToken)
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
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// close Discord session
	err := discordSession.Close()
	if err != nil {
		fmt.Println("Error: error closing session,", sessionError)
		return
	}
}