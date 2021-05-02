package Subarrays_with_Product_Less_than_a_Target

import (
	"reflect"
	"testing"
)

func Test_findSubarrays(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[]int{2, 5, 3, 10}, 30}, [][]int{{2}, {5}, {2, 5}, {3}, {5, 3}, {10}}},
		{"case2", args{[]int{8, 2, 6, 5}, 50}, [][]int{{8}, {2}, {8, 2}, {6}, {2, 6}, {5}, {6, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubarrays(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
