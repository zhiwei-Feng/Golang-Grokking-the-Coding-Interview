package Unique_Generalized_Abbreviations

import (
	"testing"
)

func Test_generateAbbreviations(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{"BAT"}, []string{"BAT", "BA1", "B1T", "B2", "1AT", "1A1", "2T", "3"}},
		{"case2", args{"code"}, []string{"code", "cod1", "co1e", "co2", "c1de", "c1d1", "c2e", "c3", "1ode", "1od1", "1o1e", "1o2",
			"2de", "2d1", "3e", "4"}},
		{"case3", args{""}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateAbbreviations(tt.args.word)
			want := tt.want
			if len(got) != len(want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				return
			}

			testMap := make(map[string]int)
			for _, v := range want {
				testMap[v] = 1
			}
			for _, v := range got {
				delete(testMap, v)
			}

			if len(testMap) != 0 {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
