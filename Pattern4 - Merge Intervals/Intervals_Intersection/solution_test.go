package Intervals_Intersection

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		arr1 []Interval
		arr2 []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"case1", args{
			[]Interval{{1, 3}, {5, 6}, {7, 9}},
			[]Interval{{2, 3}, {5, 7}}},
			[]Interval{{2, 3}, {5, 6}, {7, 7}}},
		{"case2", args{
			arr1: []Interval{{1, 3}, {5, 7}, {9, 12}},
			arr2: []Interval{{5, 10}},
		}, []Interval{{5, 7}, {9, 10}}},
		{"case3", args{
			arr1: []Interval{{3, 6}, {7, 10}},
			arr2: []Interval{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		}, []Interval{{3, 4}, {5, 6}, {7, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
