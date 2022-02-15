package rosbot

import (
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
