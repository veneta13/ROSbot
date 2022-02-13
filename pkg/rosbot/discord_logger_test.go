package rosbot

import "testing"

func Test_logger(t *testing.T) {
	type args struct {
		messageCode int
		arguments   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test for case 1", args{1, []string{""}}, message1},
		{"Test for case 2", args{2, []string{""}}, message2},
		{"Test for case 3", args{3, []string{"user"}}, "<@!user>" + message3},
		{"Test for case 4", args{4, []string{""}}, message4},
		{"Test for case 5", args{5, []string{""}}, message5},
		{"Test for case 6", args{6, []string{"user"}}, "Hello user :wave:"},
		{"Test for case 7", args{7, []string{"name"}}, "name" + message7},
		{"Test for case 8", args{8, []string{"name"}}, message8 + "name"},
		{"Test for case 9", args{9, []string{""}}, message9},
		{"Test for case 10", args{10, []string{""}}, message10},
		{"Test for case 11", args{11, []string{""}}, message11},
		{"Test for case 12", args{12, []string{""}}, message12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logger(tt.args.messageCode, tt.args.arguments...); got != tt.want {
				t.Errorf("logger() = %v, want %v", got, tt.want)
			}
		})
	}
}
