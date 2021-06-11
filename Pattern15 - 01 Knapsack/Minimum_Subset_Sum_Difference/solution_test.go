package minimumsubsetsumdifference

import (
	"testing"
)

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 9}}, 3},
		{"case2", args{[]int{1, 2, 7, 1, 5}}, 0},
		{"case3", args{[]int{1, 3, 100, 4}}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
