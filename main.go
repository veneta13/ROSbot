package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:8888/callback"

var (
	DiscordToken string
	discordSession *discordgo.Session
	client *spotify.Client
	user *spotify.PrivateUser

	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI),
							spotifyauth.WithScopes(
								spotifyauth.ScopeUserReadPrivate,
								spotifyauth.ScopePlaylistModifyPublic,
								spotifyauth.ScopePlaylistModifyPrivate,
								spotifyauth.ScopeUserLibraryModify,
								spotifyauth.ScopeUserLibraryRead,
								spotifyauth.ScopeUserTopRead))
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