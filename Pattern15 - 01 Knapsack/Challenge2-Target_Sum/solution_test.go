package Challenge2_Target_Sum

import (
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 1, 2, 3}, 1}, 3},
		{"case2", args{[]int{1, 2, 7, 1}, 9}, 2},
		{"case3", args{[]int{1, 2, 3, 4}, 20}, 0},
		{"case4", args{[]int{1, 2, 3, 4}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
