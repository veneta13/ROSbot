package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
)

func createPlaylist(mood string) (*spotify.FullPlaylist, error) {
	fmt.Println("Log: Creating playlist")
	playlist, err := client.CreatePlaylistForUser(
		context.Background(),
		user.ID,
		user.DisplayName + "'s super cool " + mood + " playlist",
		"Playlist created by ROSbot",
		false,
		false)

	if err != nil {
		fmt.Println("Error creating playlist")
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

func readImage(imagePath string) (io.Reader, error) {
	file, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error loading playlist image")
	}
	reader := bytes.NewReader(file)
	return reader, err
}

func makeCompletePlaylist(mood string, coverImage string) (playlist *spotify.FullPlaylist, err error) {
	playlist, err = createPlaylist(mood)
	if err != nil {
		return nil, err
	}

	tracks, err := getTracks(mood)
	if err != nil {
		fmt.Println("Error finding tracks")
		return nil, err
	}

	for _,track := range tracks {
		_, err := client.AddTracksToPlaylist(context.Background(), playlist.ID, track)
		if err != nil {
			fmt.Println("Error adding tracks to playlist")
			return nil, err
		}
	}

	image, err := readImage(coverImage)
	if err != nil {
		return nil, err
	}

	err = client.SetPlaylistImage(context.Background(), playlist.ID, image)
	if err != nil {
		fmt.Println("Error setting playlist image")
		return nil, err
	}

	return playlist, err
}

func getPlaylistByMood (message string) (*spotify.FullPlaylist, error) {
	if strings.Contains(message, "happy"){
		return makeCompletePlaylist("happy", playlistCoverFile)
	}
	if strings.Contains(message, "sad"){
		return makeCompletePlaylist("sad", playlistCoverFile)
	}
	if strings.Contains(message, "relaxed"){
		return makeCompletePlaylist("relaxing", playlistCoverFile)
	}
	if strings.Contains(message, "party"){
		return makeCompletePlaylist("party", playlistCoverFile)
	}
	if strings.Contains(message, "focused"){
		return makeCompletePlaylist("focus", playlistCoverFile)
	}
	if strings.Contains(message, "romantic"){
		return makeCompletePlaylist("romance", playlistCoverFile)
	}
	if strings.Contains(message, "holiday"){
		return makeCompletePlaylist("holiday", playlistCoverFile)
	}
	if strings.Contains(message, "travel"){
		return makeCompletePlaylist("road", playlistCoverFile)
	}
	if strings.Contains(message, "motivated"){
		return makeCompletePlaylist("motivational", playlistCoverFile)
	}
	if strings.Contains(message, "sleepy"){
		return makeCompletePlaylist("sleep", playlistCoverFile)
	}
	return nil, errors.New("error: mood not found")
}
