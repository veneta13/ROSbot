package rosbot

import (
	"context"
	"crypto/tls"
	"errors"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	// ScopeImageUpload seeks permission to upload images to Spotify on your behalf.
	ScopeImageUpload = "ugc-image-upload"
	// ScopePlaylistReadPrivate seeks permission to read
	// a user's private playlists.
	ScopePlaylistReadPrivate = "playlist-read-private"
	// ScopePlaylistModifyPublic seeks write access
	// to a user's public playlists.
	ScopePlaylistModifyPublic = "playlist-modify-public"
	// ScopePlaylistModifyPrivate seeks write access to
	// a user's private playlists.
	ScopePlaylistModifyPrivate = "playlist-modify-private"
	// ScopePlaylistReadCollaborative seeks permission to
	// access a user's collaborative playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	// ScopeUserFollowModify seeks write/delete access to
	// the list of artists and other users that a user follows.
	ScopeUserFollowModify = "user-follow-modify"
	// ScopeUserFollowRead seeks read access to the list of
	// artists and other users that a user follows.
	ScopeUserFollowRead = "user-follow-read"
	// ScopeUserLibraryModify seeks write/delete access to a
	// user's "Your Music" library.
	ScopeUserLibraryModify = "user-library-modify"
	// ScopeUserLibraryRead seeks read access to a user's "Your Music" library.
	ScopeUserLibraryRead = "user-library-read"
	// ScopeUserReadPrivate seeks read access to a user's
	// subsription details (type of user account).
	ScopeUserReadPrivate = "user-read-private"
	// ScopeUserReadEmail seeks read access to a user's email address.
	ScopeUserReadEmail = "user-read-email"
	// ScopeUserReadCurrentlyPlaying seeks read access to a user's currently playing track
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"
	// ScopeUserReadPlaybackState seeks read access to the user's current playback state
	ScopeUserReadPlaybackState = "user-read-playback-state"
	// ScopeUserModifyPlaybackState seeks write access to the user's current playback state
	ScopeUserModifyPlaybackState = "user-modify-playback-state"
	// ScopeUserReadRecentlyPlayed allows access to a user's recently-played songs
	ScopeUserReadRecentlyPlayed = "user-read-recently-played"
	// ScopeUserTopRead seeks read access to a user's top tracks and artists
	ScopeUserTopRead = "user-top-read"
	// ScopeStreaming seeks permission to play music and control playback on your other devices.
	ScopeStreaming = "streaming"
)

type Authenticator struct {
	config *oauth2.Config
}

type AuthenticatorOption func(a *Authenticator)

func WithClientID(id string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientID = id
	}
}

func WithClientSecret(secret string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientSecret = secret
	}
}

func WithScopes(scopes ...string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.Scopes = scopes
	}
}

func WithRedirectURL(url string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.RedirectURL = url
	}
}

func New(opts ...AuthenticatorOption) *Authenticator {
	cfg := &oauth2.Config{
		ClientID:     "place",
		ClientSecret: "holder",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
	}

	a := &Authenticator{
		config: cfg,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func contextWithHTTPClient(ctx context.Context) context.Context {
	tr := &http.Transport{
		TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
	}

	return context.WithValue(ctx, oauth2.HTTPClient, &http.Client{Transport: tr})
}

var ShowDialog = oauth2.SetAuthURLParam("show_dialog", "true")

func (a Authenticator) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return a.config.AuthCodeURL(state, opts...)
}

func (a Authenticator) Token(
	ctx context.Context, state string, r *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, errors.New("spotify: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("spotify: didn't get access code")
	}
	actualState := values.Get("state")
	if actualState != state {
		return nil, errors.New("spotify: redirect state parameter doesn't match")
	}
	return a.config.Exchange(contextWithHTTPClient(ctx), code, opts...)
}

func (a Authenticator) Exchange(
	ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return a.config.Exchange(contextWithHTTPClient(ctx), code, opts...)
}

func (a Authenticator) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.config.Client(contextWithHTTPClient(ctx), token)
}
