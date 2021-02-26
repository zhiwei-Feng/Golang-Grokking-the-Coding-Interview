package Challenge1_Permutation_in_a_String

import "testing"

func Test_findPermutation(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"oidbcaf", "abc"}, true},
		{"case2", args{"odicf", "dc"}, false},
		{"case3", args{"bcdxabcdy", "bcdyabcdx"}, true},
		{"case4", args{"aaacb", "abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPermutation(tt.args.str, tt.args.pattern); got != tt.want {
				t.Errorf("findPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
