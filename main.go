package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
)

const pattern = "/callback"
const redirectUri = "http://localhost:8888/callback"
const playlistCoverURL = "https://raw.githubusercontent.com/veneta13/ROSbot/master/assets/playlist.png"
const playlistCoverFile = "./assets/playlist.png"
const authURL = "https://accounts.spotify.com/authorize"
const tokenURL = "https://accounts.spotify.com/api/token"

var (
	DiscordToken string
	port string
	clientSecret string
	clientID string
	discordSession *discordgo.Session
	client *spotify.Client
	user *spotify.PrivateUser
	auth = New(WithRedirectURL(redirectUri))
	ch    = make(chan *spotify.Client)
	state = "myState"
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "ROSbot Token")
	flag.StringVar(&port, "port", ":8888", "The port the application is running on")
	flag.Parse()
}

func main() {

	startServer()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	stopServer()
}