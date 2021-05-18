package twosinglenumbers

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 4, 2, 1, 3, 5, 6, 2, 3, 5}}, []int{4, 6}},
		{"case2", args{[]int{2, 1, 3, 2}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.nums)
			m := make(map[int]int)
			for _, v := range got {
				m[v] = 1
			}
			for _, v := range tt.want {
				delete(m, v)
			}
			if len(m) != 0 {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
