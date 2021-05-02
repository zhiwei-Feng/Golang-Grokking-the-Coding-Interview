package Challenge3_Smallest_Window_containing_Substring

import "testing"

func Test_findSubstring(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"aabdec", "abc"}, "abdec"},
		{"case2", args{"abdbca", "abc"}, "bca"},
		{"case3", args{"adcad", "abc"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.str, tt.args.pattern); got != tt.want {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
