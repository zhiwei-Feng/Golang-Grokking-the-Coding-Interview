package Topological_Sort

import (
	"reflect"
	"testing"
)

func Test_topologicSort(t *testing.T) {
	type args struct {
		vertices int
		edges    [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{4, [][]int{{3, 2}, {3, 0}, {2, 0}, {2, 1}}}, []int{3, 2, 0, 1}},
		{"case2", args{0, [][]int{}}, []int{}},
		{"case3", args{5, [][]int{{4, 2}, {4, 3}, {2, 0}, {2, 1}, {3, 1}}}, []int{4, 2, 3, 0, 1}},
		{"case4", args{5, [][]int{{4, 2}, {4, 3}, {2, 0}, {2, 1}, {3, 1}, {1, 3}}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topologicSort(tt.args.vertices, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologicSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
