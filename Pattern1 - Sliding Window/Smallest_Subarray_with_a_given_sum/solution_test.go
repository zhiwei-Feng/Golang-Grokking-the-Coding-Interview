package Smallest_Subarray_with_a_given_sum

import "testing"

func Test_findMinSubArray(t *testing.T) {
	type args struct {
		S   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{7, []int{2, 1, 5, 2, 3, 2}}, want: 2},
		{name: "case2", args: args{7, []int{2, 1, 5, 2, 8}}, want: 1},
		{name: "case3", args: args{8, []int{3, 4, 1, 1, 6}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinSubArray(tt.args.S, tt.args.arr); got != tt.want {
				t.Errorf("findMinSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
