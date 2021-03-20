package Merge_Intervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"case1", args{[]Interval{{1, 4}, {2, 5}, {7, 9}}}, []Interval{{1, 5}, {7, 9}}},
		{"case2", args{[]Interval{{6, 7}, {2, 4}, {5, 9}}}, []Interval{{2, 4}, {5, 9}}},
		{"case3", args{[]Interval{{1, 4}, {2, 6}, {3, 5}}}, []Interval{{1, 6}}},
		{"case4", args{[]Interval{{1, 4}}}, []Interval{{1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
