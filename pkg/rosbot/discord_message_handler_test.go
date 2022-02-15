package rosbot

import (
	"github.com/zmb3/spotify/v2"
	"reflect"
	"testing"
)

func Test_getClientCredentials(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name             string
		args             args
		wantClientID     string
		wantClientSecret string
	}{
		{"Log in", args{"!log-in ID=id SECRET=secret"}, "id", "secret"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClientID, gotClientSecret := getClientCredentials(tt.args.message)
			if gotClientID != tt.wantClientID {
				t.Errorf("getClientCredentials() gotClientID = %v, want %v", gotClientID, tt.wantClientID)
			}
			if gotClientSecret != tt.wantClientSecret {
				t.Errorf("getClientCredentials() gotClientSecret = %v, want %v", gotClientSecret, tt.wantClientSecret)
			}
		})
	}
}

func Test_getClient(t *testing.T) {
	type args struct {
		discordID string
	}
	tests := []struct {
		name string
		args args
		want *spotify.Client
	}{
		{name: "Client not found",
		 args: args {discordID: "myDiscordID"},
		 want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(tt.args.discordID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}