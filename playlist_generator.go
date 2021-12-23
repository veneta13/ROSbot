package main

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
)

func createPlaylist(mood string) (*spotify.FullPlaylist, error) {
	fmt.Println("Log: Creating playlist")

	playlist, err := client.CreatePlaylistForUser(context.Background(),
												  user.ID,
												  user.DisplayName + "'s super cool " + mood + " playlist",
												  "Playlist created by ROSbot",
												  false,
												  false)

	if err == nil {
		fmt.Println("Log: Created playlist")
	}

	return playlist, err
}
