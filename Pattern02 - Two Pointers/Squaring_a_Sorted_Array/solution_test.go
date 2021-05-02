package Squaring_a_Sorted_Array

import (
	"reflect"
	"testing"
)

func Test_makeSquares(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{-2, -1, 0, 2, 3}}, []int{0, 1, 4, 4, 9}},
		{"case2", args{[]int{-3, -1, 0, 1, 2}}, []int{0, 1, 1, 4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSquares(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
