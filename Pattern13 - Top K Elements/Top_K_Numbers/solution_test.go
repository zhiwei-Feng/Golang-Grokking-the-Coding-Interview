package topknumbers

import (
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{3, 1, 5, 12, 2, 11}, 3}, []int{5, 12, 11}},
		{"case2", args{[]int{5, 12, 11, -1, 12}, 3}, []int{12, 11, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLargestNumbers(tt.args.nums, tt.args.k)
			m := make(map[int]int)
			for _, v := range got {
				m[v] = v
			}
			for _, v := range tt.want {
				delete(m, v)
			}
			if len(m) != 0 {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
