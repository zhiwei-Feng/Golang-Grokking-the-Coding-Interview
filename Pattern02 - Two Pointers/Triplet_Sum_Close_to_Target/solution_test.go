package Triplet_Sum_Close_to_Target

import "testing"

func Test_searchTriplet(t *testing.T) {
	type args struct {
		arr       []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{-2, 0, 1, 2}, 2}, 1},
		{"case2", args{[]int{-3, -1, 1, 2}, 1}, 0},
		{"case3", args{[]int{1, 0, 1, 1}, 100}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchTriplet(tt.args.arr, tt.args.targetSum); got != tt.want {
				t.Errorf("searchTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
