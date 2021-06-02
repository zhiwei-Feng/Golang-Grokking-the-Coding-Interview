package kpairswithlargestsums

import (
	"reflect"
	"testing"
)

func Test_kLargestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[]int{9, 8, 2}, []int{6, 3, 1}, 3}, [][]int{{9, 3}, {9, 6}, {8, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLargestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kLargestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
