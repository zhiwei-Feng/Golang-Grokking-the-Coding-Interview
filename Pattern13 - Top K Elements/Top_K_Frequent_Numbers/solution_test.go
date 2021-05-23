package topkfrequentnumbers

import (
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			m := make(map[int]int)
			for _, v := range got {
				m[v]++
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
