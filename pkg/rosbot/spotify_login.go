package rosbot

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"net/http"
	"strconv"
)

// SpotifyLogin Spotify authentication
func SpotifyLogin() string {
	commandLineLogger(14)

	_, loginLink := AuthUser()

	return loginLink
}

// AuthUser Authenticate user
func AuthUser() (context.Context, string) {
	http.HandleFunc(projectProperties.pattern, completeAuth)

	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(projectProperties.port), nil)
		if err != nil {
			commandLineLogger(2)

			return
		}
	}()

	url := auth.AuthURL(projectProperties.state)
	loginLink := "Log in via Spotify here :point_right: " + url

	return context.Background(), loginLink
}

// complete authentication
func completeAuth(writer http.ResponseWriter, reader *http.Request) {
	tok, err := auth.Token(reader.Context(), projectProperties.state, reader)

	if err != nil {
		http.Error(writer, "Error: Error getting Spotify token", http.StatusForbidden)
		commandLineLogger(15)
	}

	client := spotify.New(auth.Client(reader.Context(), tok))
	_, err = fmt.Fprintf(writer, "Log: Login Completed")

	if err != nil {
		http.Error(writer, "Error: Cannot log in", http.StatusForbidden)
		commandLineLogger(16)
	}

	ch <- client

	commandLineLogger(17)
}
