package Triplets_with_Smaller_Sum

import "testing"

func Test_searchTriplets(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{-1, 0, 2, 3}, 3}, 2},
		{"case2", args{[]int{-1, 4, 2, 1, 3}, 5}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchTriplets(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("searchTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
