package Challenge4_Words_Concatenation

import (
	"reflect"
	"testing"
)

func Test_findWordConcatenation(t *testing.T) {
	type args struct {
		str   string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{"catfoxcat", []string{"cat", "fox"}}, []int{0, 3}},
		{"case2", args{"catcatfoxfox", []string{"cat", "fox"}}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordConcatenation(tt.args.str, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
