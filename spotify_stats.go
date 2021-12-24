package main

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"strings"
)

func getArtistStats(time string) (*spotify.FullArtistPage, error) {
	artists, err := client.CurrentUsersTopArtists(
		context.Background(),
		spotify.Limit(5),
		spotify.Timerange(spotify.Range(time)))
	if err != nil {
		fmt.Println("Error getting artist stats")
	}
	return artists, err
}

func getTrackStats(time string) (*spotify.FullTrackPage, error) {
	tracks, err := client.CurrentUsersTopTracks(
		context.Background(),
		spotify.Limit(10),
		spotify.Timerange(spotify.Range(time)))
	if err != nil {
		fmt.Println("Error getting track stats")
	}
	return tracks, err
}

func getStatsType(message string) (statsType int, statsTime string) {
	if strings.Contains(message, "artist") {
		statsType = 1
	} else if strings.Contains(message, "track"){
		statsType = 2
	} else if strings.Contains(message, "full") {
		statsType = 3
	}

	if strings.Contains(message, "last month") {
		statsTime = "short_term"
	} else if strings.Contains(message, "6 months"){
		statsTime = "medium_term"
	} else if strings.Contains(message, "all time") {
		statsTime = "long_term"
	}
	
	return
}

func getStats(statsType int, statsTime string) (tracks *spotify.FullTrackPage, artists *spotify.FullArtistPage, err error) {
	switch statsType {
	case 1:
		artists, err = getArtistStats(statsTime)
		return
	case 2:
		tracks, err = getTrackStats(statsTime)
		return
	case 3:
		artists, err = getArtistStats(statsTime)
		if err != nil {
			return
		}
		tracks, err = getTrackStats(statsTime)
		return
	}
	return
}

func makeTrackList(tracks *spotify.FullTrackPage) (trackList []spotify.FullTrack) {
	for _, track := range tracks.Tracks {
		trackList = append(trackList, track)
	}
	return
}

func makeArtistList(artists *spotify.FullArtistPage) (artistList []spotify.FullArtist) {
	for _, artist := range artists.Artists {
		artistList = append(artistList, artist)
	}
	return
}
