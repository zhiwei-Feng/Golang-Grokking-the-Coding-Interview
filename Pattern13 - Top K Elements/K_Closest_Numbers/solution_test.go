package kclosestnumbers

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, 4, 3}, []int{1, 2, 3, 4}},
		{"case2", args{[]int{1, 2, 3, 4, 5}, 4, -1}, []int{1, 2, 3, 4}},
		{"case3", args{[]int{5, 6, 7, 8, 9}, 3, 7}, []int{6, 7, 8}},
		{"case4", args{[]int{2, 4, 5, 6, 9}, 3, 6}, []int{4, 5, 6}},
		{"case5", args{[]int{2, 4, 5, 6, 9}, 3, 10}, []int{5, 6, 9}},
		{"case6", args{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5}, []int{3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
