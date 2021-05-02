package Longest_Substring_with_K_Distinct_Characters

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
		{"case1", args{"araaci", 2}, 4},
		{"case2", args{"araaci", 1}, 2},
		{"case3", args{"cbbebi", 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
