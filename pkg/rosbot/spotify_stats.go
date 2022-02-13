package rosbot

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"strings"
)

// Get artist statistics in set timeframe for current client.
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

// Get track statistics in set timeframe for current client.
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

// Extract the statistics type (artist/track/full) and statistics period
// (last month/6 months/all time) from user request message.
func getStatsType(message string) (statsType int, statsTime string) {
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

// Get statistics from Spotify API for current user according
// to request parameters.
func getStats(statsType int, statsTime string, client *spotify.Client) (
	tracks *spotify.FullTrackPage,
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

// Transform Spotify response from FullTrackPage to array of FullTrack values.
func makeTrackList(tracks *spotify.FullTrackPage) (trackList []spotify.FullTrack) {
	trackList = append(trackList, tracks.Tracks...)

	return
}

// Transform Spotify response from FullArtistPage to array of FullArtist values.
func makeArtistList(artists *spotify.FullArtistPage) (artistList []spotify.FullArtist) {
	artistList = append(artistList, artists.Artists...)

	return
}
