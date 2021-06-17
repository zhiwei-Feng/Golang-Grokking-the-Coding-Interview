package Challenge2

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{5, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}}}, []int{1, 2}},
		{"case2", args{4, [][]int{{0, 1}, {0, 2}, {2, 3}}}, []int{0, 2}},
		{"case3", args{4, [][]int{{0, 1}, {1, 2}, {1, 3}}}, []int{1}},
		{"case4", args{6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}}, []int{3, 4}},
		{"case5", args{1, [][]int{}}, []int{0}},
		{"case6", args{2, [][]int{{0, 1}}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
