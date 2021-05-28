package challenge1

import "testing"

func Test_rearrangeString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"aabbcc", 3}, "abcabc"},
		{"case2", args{"aaabc", 3}, ""},
		{"case3", args{"aaadbbcc", 2}, "abacabcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeString(tt.args.s, tt.args.k); len(got) != len(tt.want) {
				t.Errorf("rearrangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
