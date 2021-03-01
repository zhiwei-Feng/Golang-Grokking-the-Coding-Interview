package Challenge2_String_Anagrams

import (
	"reflect"
	"testing"
)

func Test_findStringAnagrams(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{"ppqp", "pq"}, []int{1, 2}},
		{"case2", args{"abbcabc", "abc"}, []int{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStringAnagrams(tt.args.str, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStringAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
