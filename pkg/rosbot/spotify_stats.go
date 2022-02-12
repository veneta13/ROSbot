package rosbot

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"strings"
)

func getArtistStats(time string, client *spotify.Client) (*spotify.FullArtistPage, error) {
	artists, err := client.CurrentUsersTopArtists(
		context.Background(),
		spotify.Limit(5),
		spotify.Timerange(spotify.Range(time)))
	if err != nil {
		commandLineLogger(25)
	}

	return artists, err
}

func getTrackStats(time string, client *spotify.Client) (*spotify.FullTrackPage, error) {
	tracks, err := client.CurrentUsersTopTracks(
		context.Background(),
		spotify.Limit(10),
		spotify.Timerange(spotify.Range(time)))
	if err != nil {
		commandLineLogger(26)
	}

	return tracks, err
}

func GetStatsType(message string) (statsType int, statsTime string) {
	switch {
	case strings.Contains(message, "artist"):
		statsType = 1
	case strings.Contains(message, "track"):
		statsType = 2
	case strings.Contains(message, "full"):
		statsType = 3
	}

	switch {
	case strings.Contains(message, "last month"):
		statsTime = "short_term"
	case strings.Contains(message, "6 months"):
		statsTime = "medium_term"
	case strings.Contains(message, "all time"):
		statsTime = "long_term"
	}
	
	return
}

func GetStats(statsType int, statsTime string, client *spotify.Client) (tracks *spotify.FullTrackPage,
	artists *spotify.FullArtistPage,
	err error) {
	switch statsType {
	case 1:
		artists, err = getArtistStats(statsTime, client)

		return
	case 2:
		tracks, err = getTrackStats(statsTime, client)

		return
	case 3:
		artists, err = getArtistStats(statsTime, client)
		if err != nil {
			return
		}

		tracks, err = getTrackStats(statsTime, client)

		return
	}

	return
}

func makeTrackList(tracks *spotify.FullTrackPage) (trackList []spotify.FullTrack) {
	trackList = append(trackList, tracks.Tracks...)

	return
}

func makeArtistList(artists *spotify.FullArtistPage) (artistList []spotify.FullArtist) {
	artistList = append(artistList, artists.Artists...)

	return
}
