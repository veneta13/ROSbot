package rosbot

import (
	"bytes"
	"context"
	"errors"
	"github.com/zmb3/spotify/v2"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
)

func createPlaylist(mood string, client *spotify.Client) (*spotify.FullPlaylist, error) {
	commandLineLogger(18)

	user, _ := client.CurrentUser(context.Background())
	playlist, err := client.CreatePlaylistForUser(
		context.Background(),
		user.ID,
		user.DisplayName+"'s super cool "+mood+" playlist",
		"Playlist created by ROSbot",
		false,
		false)

	if err != nil {
		commandLineLogger(19)
	}

	return playlist, err
}

func getTracks(mood string, client *spotify.Client) ([]spotify.ID, error) {
	var trackIDs []spotify.ID

	result, err := client.Search(context.Background(), mood, spotify.SearchTypePlaylist)
	if err != nil {
		commandLineLogger(20)
	}

	for _, currentPlaylist := range result.Playlists.Playlists {
		playlistTracks, err := client.GetPlaylistTracks(context.Background(), currentPlaylist.ID)
		if err != nil {
			return nil, err
		}

		for _, track := range playlistTracks.Tracks {
			if rand.Int31()%100 == 10 {
				trackIDs = append(trackIDs, track.Track.ID)
			}
		}
	}

	return trackIDs, err
}

func readImage(imagePath string) (io.Reader, error) {
	file, err := ioutil.ReadFile(imagePath)
	if err != nil {
		commandLineLogger(21)
	}

	reader := bytes.NewReader(file)

	return reader, err
}

func makeCompletePlaylist(mood string,
	coverImage string,
	client *spotify.Client) (playlist *spotify.FullPlaylist, err error) {
	playlist, err = createPlaylist(mood, client)
	if err != nil {
		return nil, err
	}

	tracks, err := getTracks(mood, client)
	if err != nil {
		commandLineLogger(22)

		return nil, err
	}

	for _, track := range tracks {
		_, err := client.AddTracksToPlaylist(context.Background(), playlist.ID, track)
		if err != nil {
			commandLineLogger(23)

			return nil, err
		}
	}

	image, err := readImage(coverImage)
	if err != nil {
		return nil, err
	}

	err = client.SetPlaylistImage(context.Background(), playlist.ID, image)
	if err != nil {
		commandLineLogger(24)

		return nil, err
	}

	return playlist, err
}

func GetPlaylistByMood(message string, client *spotify.Client) (*spotify.FullPlaylist, error) {
	if strings.Contains(message, "happy") {
		return makeCompletePlaylist("happy", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "sad") {
		return makeCompletePlaylist("sad", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "relaxed") {
		return makeCompletePlaylist("relaxing", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "party") {
		return makeCompletePlaylist("party", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "focused") {
		return makeCompletePlaylist("focus", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "romantic") {
		return makeCompletePlaylist("romance", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "holiday") {
		return makeCompletePlaylist("holiday", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "travel") {
		return makeCompletePlaylist("road", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "motivated") {
		return makeCompletePlaylist("motivational", projectProperties.playlistCoverFile, client)
	}

	if strings.Contains(message, "sleepy") {
		return makeCompletePlaylist("sleep", projectProperties.playlistCoverFile, client)
	}

	return nil, errors.New("error: mood not found")
}
