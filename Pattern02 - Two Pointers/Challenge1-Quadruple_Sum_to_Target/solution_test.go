package Challenge1_Quadruple_Sum_to_Target

import (
	"reflect"
	"testing"
)

func Test_searchQuadruplets(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][4]int
	}{
		{"case1", args{[]int{4, 1, 2, -1, 1, -3}, 1}, [][4]int{{-3, -1, 1, 4}, {-3, 1, 1, 2}}},
		{"case2", args{[]int{2, 0, -1, 1, -2, 2}, 2}, [][4]int{{-2, 0, 2, 2}, {-1, 0, 1, 2}}},
		{"case3", args{[]int{2, 0, -1, 1, -2, 2, 3, 3, -2}, 2}, [][4]int{{-2, -2, 3, 3}, {-2, -1, 2, 3}, {-2, 0, 1, 3}, {-2, 0, 2, 2}, {-1, 0, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchQuadruplets(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
