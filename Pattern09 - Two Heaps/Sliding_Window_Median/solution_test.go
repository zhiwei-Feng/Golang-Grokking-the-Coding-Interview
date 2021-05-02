package Sliding_Window_Median

import (
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"case1", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []float64{1.0, -1.0, -1.0, 3.0, 5.0, 6.0}},
		{"case2", args{[]int{}, 3}, []float64{}},
		{"case3", args{[]int{2147483647, 1, 2, 3, 4, 5, 6, 7, 2147483647}, 2}, []float64{1073741824.00000, 1.50000, 2.50000, 3.50000, 4.50000, 5.50000, 6.50000, 1073741827.00000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
