package rosbot

import "golang.org/x/oauth2"

// The code in this file is taken from https://github.com/zmb3/spotify/blob/master/auth/auth.go.
// There are changes made to the New function in order for it not to read from the
// environment variables but use placeholder values instead.

import (
	"context"
	"crypto/tls"
	"errors"
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

// Authenticator provides convenience functions for implementing the OAuth2 flow.
// You should always use `New` to make them.
type Authenticator struct {
	config *oauth2.Config
}

// AuthenticatorOption provides options for the authenticator.
type AuthenticatorOption func(a *Authenticator)

// WithClientID allows a client ID to be specified. Without this the value of the SPOTIFY_ID environment
// variable will be used.
func WithClientID(id string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientID = id
	}
}

// WithClientSecret allows a client secret to be specified. Without this the value of the SPOTIFY_SECRET environment
// variable will be used.
func WithClientSecret(secret string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientSecret = secret
	}
}

// WithScopes configures the oauth scopes that the client should request.
func WithScopes(scopes ...string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.Scopes = scopes
	}
}

// WithRedirectURL configures a redirect url for oauth flows. It must exactly match one of the
// URLs specified in your Spotify developer account.
func WithRedirectURL(url string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.RedirectURL = url
	}
}

// New creates an authenticator which is used to implement the OAuth2 authorization flow.
// Here the default values of ClientID and ClientSecret are placeholder values,
// unlike in the zmb3/spotify package where they are extracted from
// the environment variables.
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

// ShowDialog forces the user to approve the app, even if they have already done so.
// Without this, users who have already approved the app are immediately redirected to the redirect uri.
var ShowDialog = oauth2.SetAuthURLParam("show_dialog", "true")

// AuthURL returns a URL to the the Spotify Accounts Service's OAuth2 endpoint.
//
// State is a token to protect the user from CSRF attacks.  You should pass the
// same state to `Token`, where it will be validated.  For more info, refer to
// http://tools.ietf.org/html/rfc6749#section-10.12.
func (a Authenticator) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return a.config.AuthCodeURL(state, opts...)
}

// Token pulls an authorization code from an HTTP request and attempts to exchange
// it for an access token.  The standard use case is to call Token from the handler
// that handles requests to your application's redirect URL.
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

// Exchange is like Token, except it allows you to manually specify the access
// code instead of pulling it out of an HTTP request.
func (a Authenticator) Exchange(
	ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return a.config.Exchange(contextWithHTTPClient(ctx), code, opts...)
}

// Client creates a *http.Client that will use the specified access token for its API requests.
// Combine this with spotify.HTTPClientOpt.
func (a Authenticator) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.config.Client(contextWithHTTPClient(ctx), token)
}
