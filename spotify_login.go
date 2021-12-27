package main

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"net/http"
)

func spotifyLogin() string {
	// Spotify authentication
	fmt.Println("spotify-login")
	_ , loginLink := AuthUser()
	return loginLink
}

func AuthUser() (context.Context, string) {

	http.HandleFunc("/callback", completeAuth)
	go func() {
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			fmt.Println("Error: Callback error")
			return
		}
	}()

	url := auth.AuthURL(state)
	loginLink := "Log in via Spotify here :point_right: " + url
	return context.Background(), loginLink
}

func completeAuth(writer http.ResponseWriter, reader *http.Request) {
	tok, err := auth.Token(reader.Context(), state, reader)
	if err != nil {
		http.Error(writer, "Error: Error getting Spotify token", http.StatusForbidden)
		fmt.Println("Error: Error getting Spotify token")
	}

	client := spotify.New(auth.Client(reader.Context(), tok))
	_, err = fmt.Fprintf(writer, "Log: Login Completed")
	if err != nil {
		http.Error(writer, "Error: Cannot log in", http.StatusForbidden)
		fmt.Println("Error: Cannot log in")
	}

	ch <- client
	fmt.Println("Log: Login successful")
}
