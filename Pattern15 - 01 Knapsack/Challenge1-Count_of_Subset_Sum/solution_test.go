package challenge1countofsubsetsum

import (
	"testing"
)

func Test_countSubsets(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 1, 2, 3}, 4}, 3},
		{"case2", args{[]int{1, 2, 7, 1, 5}, 9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubsets(tt.args.nums, tt.args.sum); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
