package Cyclic_Sort

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[]int{3, 1, 5, 4, 2}}},
		{"case2", args{[]int{2, 6, 4, 3, 1, 5}}},
		{"case3", args{[]int{1, 5, 6, 4, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.nums)
			for i, v := range tt.args.nums {
				if i+1 != v {
					t.Errorf("want %d, but got %d", i, v)
					break
				}
			}
			fmt.Println(tt.args.nums)
		})
	}
}
