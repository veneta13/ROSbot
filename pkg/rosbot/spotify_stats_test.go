package rosbot

import (
	"testing"
)

func Test_getStatsType(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name          string
		args          args
		wantStatsType int
		wantStatsTime string
	}{
		{"Artist, last month", args{"!get-stats artist last month"}, 1, "short_term"},
		{"Artist, 6 months", args{"!get-stats artist 6 months"}, 1, "medium_term"},
		{"Artist, all time", args{"!get-stats artist all time"}, 1, "long_term"},
		{"Track, last month", args{"!get-stats track last month"}, 2, "short_term"},
		{"Track, 6 months", args{"!get-stats track 6 months"}, 2, "medium_term"},
		{"Track, all time", args{"!get-stats track all time"}, 2, "long_term"},
		{"Full, last month", args{"!get-stats full last month"}, 3, "short_term"},
		{"Full, 6 months", args{"!get-stats full 6 months"}, 3, "medium_term"},
		{"Full, all time", args{"!get-stats full all time"}, 3, "long_term"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatsType, gotStatsTime := getStatsType(tt.args.message)
			if gotStatsType != tt.wantStatsType {
				t.Errorf("getStatsType() gotStatsType = %v, want %v", gotStatsType, tt.wantStatsType)
			}
			if gotStatsTime != tt.wantStatsTime {
				t.Errorf("getStatsType() gotStatsTime = %v, want %v", gotStatsTime, tt.wantStatsTime)
			}
		})
	}
}
