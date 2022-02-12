package rosbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
	"strconv"
)

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

var ch chan *spotify.Client
var users map[string]*spotify.Client
var auth *Authenticator
var discordSession *discordgo.Session

// StartServer Starts the Discord server session.
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

	// messageCreate - callback for MessageCreate events
	discordSession.AddHandler(messageCreate)

	// receiving message events
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	// open websocket connection
	sessionError = discordSession.Open()
	if sessionError != nil {
		fmt.Println(sessionError)
		commandLineLogger(3)

		return
	}

	commandLineLogger(4)
}

func StopServer() {
	if err := discordSession.Close(); err != nil {
		commandLineLogger(5)
	}
}
