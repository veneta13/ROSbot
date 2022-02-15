package rosbot

import "testing"

func Test_commandLineLogger(t *testing.T) {
	type args struct {
		messageCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Output 1", args{1}},
		{"Output 2", args{2}},
		{"Output 3", args{3}},
		{"Output 4", args{4}},
		{"Output 5", args{5}},
		{"Output 6", args{6}},
		{"Output 7", args{7}},
		{"Output 8", args{8}},
		{"Output 9", args{9}},
		{"Output 10", args{10}},
		{"Output 11", args{11}},
		{"Output 12", args{12}},
		{"Output 13", args{13}},
		{"Output 14", args{14}},
		{"Output 15", args{15}},
		{"Output 16", args{16}},
		{"Output 17", args{17}},
		{"Output 18", args{18}},
		{"Output 19", args{19}},
		{"Output 20", args{20}},
		{"Output 21", args{21}},
		{"Output 22", args{22}},
		{"Output 23", args{23}},
		{"Output 24", args{24}},
		{"Output 25", args{25}},
		{"Output 26", args{26}},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(commandLineMessages[i])
		})
	}
}
