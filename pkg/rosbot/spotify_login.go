package rosbot

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"net/http"
	"strconv"
)

// SpotifyLogin returns a Spotify login link for the current user.
func SpotifyLogin() string {
	commandLineLogger(14)

	_, loginLink := AuthUser()

	return loginLink
}

// AuthUser opens a port for listening to Spotify login.
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
	loginLink := logger(2) + url

	return context.Background(), loginLink
}

// Complete Spotify token authentication and print message on the browser port.
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
