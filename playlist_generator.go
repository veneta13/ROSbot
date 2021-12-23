package main

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"math/rand"
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

func getTracks(mood string) ([]spotify.ID, error) {
	result, err := client.Search(context.Background(), mood, spotify.SearchTypePlaylist)
	if err != nil {
		fmt.Println("Log: Cannot find information for keyword")
	}

	var trackIDs []spotify.ID

	for _, currentPlaylist := range result.Playlists.Playlists{
		playlistTracks, err := client.GetPlaylistTracks(context.Background(), currentPlaylist.ID)
		if err != nil {
			return nil, err
		}

		// randomize chances to add a track
		for _, track := range playlistTracks.Tracks {
			if rand.Int31() % 100  == 10 {
				trackIDs = append(trackIDs, track.Track.ID)
			}
		}
	}

	return trackIDs, err
}
