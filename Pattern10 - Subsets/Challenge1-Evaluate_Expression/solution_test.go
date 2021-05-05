package Challenge1_Evaluate_Expression

import (
	"testing"
)

func Test_diffWaysToEvaluateExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{"1+2*3"}, []int{7, 9}},
		{"case2", args{"2*3-4-5"}, []int{8, -12, 7, -7, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToEvaluateExpression(tt.args.input)
			want := tt.want
			if len(got) != len(want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				return
			}

			testMap := make(map[int]int)
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
