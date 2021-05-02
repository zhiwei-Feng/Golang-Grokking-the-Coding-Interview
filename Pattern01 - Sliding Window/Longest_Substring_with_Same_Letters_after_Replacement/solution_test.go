package Longest_Substring_with_Same_Letters_after_Replacement

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"aabccbb", 2}, 5},
		{"case2", args{"abbcb", 1}, 4},
		{"case3", args{"abccde", 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
