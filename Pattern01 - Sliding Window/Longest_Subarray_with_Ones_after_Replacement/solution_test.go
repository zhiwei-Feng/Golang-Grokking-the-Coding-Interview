package Longest_Subarray_with_Ones_after_Replacement

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2}, 6},
		{"case2", args{[]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
