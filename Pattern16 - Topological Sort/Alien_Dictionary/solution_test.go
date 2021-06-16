package Alien_Dictionary

import "testing"

func Test_findOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{[]string{"ba", "bc", "ac", "cab"}}, "bac"},
		{"case2", args{[]string{"cab", "aaa", "aab"}}, "cab"},
		{"case3", args{[]string{"ywx", "wz", "xww", "xz", "zyy", "zwz"}}, "ywxz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.words); got != tt.want {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
