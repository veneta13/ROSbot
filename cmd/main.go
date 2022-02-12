package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zmb3/spotify/v2"
	"myproject.com/module/pkg/rosbot"
)

var (
	discordToken string
	port int
)

func init() {
	flag.StringVar(&discordToken, "t", "", "ROSbot Token")
	flag.IntVar(&port, "port", 8888, "The port the application is running on")
	flag.Parse()
}

func main() {
	ch := make(chan *spotify.Client)
	rosbot.StartServer(discordToken, port, ch)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	rosbot.StopServer()
}