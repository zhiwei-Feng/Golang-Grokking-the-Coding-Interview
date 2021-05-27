package rearrangestring

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"aappp"}, "papap"},
		{"case2", args{"aapa"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
