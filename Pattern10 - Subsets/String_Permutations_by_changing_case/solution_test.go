package String_Permutations_by_changing_case

import (
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{"a1b2"}, []string{"a1b2", "A1b2", "a1B2", "A1B2"}},
		{"case2", args{"ab7c"}, []string{"ab7c", "Ab7c", "aB7c", "AB7c", "ab7C", "Ab7C", "aB7C", "AB7C"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCasePermutation(tt.args.S)
			want := tt.want
			if len(got) != len(want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
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
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
