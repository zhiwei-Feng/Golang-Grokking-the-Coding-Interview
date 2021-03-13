package Challenge2_Comparing_Strings_containing_Backspaces

import "testing"

func Test_compare(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"xy#z", "xzz#"}, true},
		{"case2", args{"xy#z", "xyz#"}, false},
		{"case3", args{"xp#", "xyz##"}, true},
		{"case4", args{"xywrrmp", "xywrrmu#p"}, true},
		{"case5", args{"", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
