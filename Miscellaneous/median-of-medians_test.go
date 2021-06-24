package Miscellaneous

import "testing"

func Test_findKthSmallestNumberUsingMedianOfMedians(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 5, 12, 2, 11, 5}, 3}, 5},
		{"case2", args{[]int{1, 5, 12, 2, 11, 5}, 4}, 5},
		{"case3", args{[]int{5, 12, 11, -1, 12}, 3}, 11},
		{"case4", args{[]int{5, 12, 11, -1, 12}, 6}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallestNumberUsingMedianOfMedians(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallestNumberUsingMedianOfMedians() = %v, want %v", got, tt.want)
			}
		})
	}
}
