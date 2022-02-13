// Package rosbot holds all bot logic for communication with the Discord and Spotify APIs.
package rosbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
	"strconv"
)

// ProjectProperties stores properties used to set up the project.
type ProjectProperties struct {
	pattern           string
	redirectURL       string
	playlistCoverURL  string
	playlistCoverFile string
	authURL           string
	tokenURL          string
	state             string
	port              int
}

var projectProperties = ProjectProperties{
	pattern:           "/callback",
	playlistCoverURL:  "https://raw.githubusercontent.com/veneta13/ROSbot/master/assets/playlist.png",
	playlistCoverFile: "./../assets/playlist.png",
	authURL:           "https://accounts.spotify.com/authorize",
	tokenURL:          "https://accounts.spotify.com/api/token",
	state:             "myState",
}

// The Spotify client channel.
var ch chan *spotify.Client

// Maps the Discord user ID to their Spotify client credentials.
var users map[string]*spotify.Client

// Spotify authenticator, storing permission information.
var auth *Authenticator

// The current Discord session.
var discordSession *discordgo.Session

// StartServer starts the Discord server session.
// Sets up additional project properties.
func StartServer(discordToken string, port int, channel chan *spotify.Client) {
	projectProperties.redirectURL = "http://localhost:" + strconv.Itoa(port) + projectProperties.pattern
	projectProperties.port = port

	users = make(map[string]*spotify.Client)

	ch = channel

	auth = New(WithRedirectURL(projectProperties.redirectURL))

	var sessionError error

	discordSession, sessionError = discordgo.New("Bot " + discordToken)

	if sessionError != nil {
		commandLineLogger(1)

		return
	}

	// Set messageCreate as callback for events.
	discordSession.AddHandler(messageCreate)

	// Receive message events.
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	// Open websocket connection.
	sessionError = discordSession.Open()
	if sessionError != nil {
		fmt.Println(sessionError)
		commandLineLogger(3)

		return
	}

	commandLineLogger(4)
}

// StopServer stops the Discord server session.
func StopServer() {
	if err := discordSession.Close(); err != nil {
		commandLineLogger(5)
	}
}
