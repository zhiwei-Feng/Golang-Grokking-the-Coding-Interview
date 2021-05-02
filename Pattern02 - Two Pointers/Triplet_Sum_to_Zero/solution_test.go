package Triplet_Sum_to_Zero

import (
	"reflect"
	"testing"
)

func Test_searchTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		{"case1", args{[]int{-3, 0, 1, 2, -1, 1, -2}}, [][3]int{{-3, 1, 2}, {-2, 0, 2},
			{-2, 1, 1}, {-1, 0, 1}}},
		{"case2", args{[]int{-5, 2, -1, -2, 3}}, [][3]int{{-5, 2, 3}, {-2, -1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchTriplets(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
