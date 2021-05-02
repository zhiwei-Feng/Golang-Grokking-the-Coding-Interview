package No_repeat_Substring

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"aabccbb"}, 3},
		{"case2", args{"abbbb"}, 2},
		{"case3", args{"abccde"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.str); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
