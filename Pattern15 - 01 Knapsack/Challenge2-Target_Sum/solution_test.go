package Challenge2_Target_Sum

import (
	"github.com/stretchr/testify/assert"
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
	}
	for _, tt := range tests {
		assert.Equal(t, findTargetSumWays(tt.args.nums, tt.args.target), tt.want)
	}
}
