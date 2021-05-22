package kclosestpointstotheorigin

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{"case2", args{[][]int{{1, 3}, {3, 4}, {2, -1}}, 2}, [][]int{{1, 3}, {2, -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
