package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
)

const RedirectUri = "http://localhost:8888/callback"
const PlaylistCoverURL = "https://raw.githubusercontent.com/veneta13/ROSbot/master/assets/playlist.png?token=AIY3LQX3UH4IYQFDD6DX43TBZ3S54"
const PlaylistCoverFile = "./assets/playlist.png"
const AuthURL = "https://accounts.spotify.com/authorize"
const TokenURL = "https://accounts.spotify.com/api/token"

var (
	DiscordToken string
	clientSecret string
	clientID string
	discordSession *discordgo.Session
	client *spotify.Client
	user *spotify.PrivateUser
	auth = New(WithRedirectURL(RedirectUri))
	ch    = make(chan *spotify.Client)
	state = "myState"
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "ROSbot Token")
	flag.Parse()
}

func main() {

	startServer()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	stopServer()
}