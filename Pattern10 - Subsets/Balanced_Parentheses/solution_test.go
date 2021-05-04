package Balanced_Parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"case2", args{1}, []string{"()"}},
		{"case3", args{4}, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
		{"case4", args{0}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}

			got := generateParenthesis(tt.args.n)
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
