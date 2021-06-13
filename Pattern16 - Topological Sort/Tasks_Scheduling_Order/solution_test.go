package Tasks_Scheduling_Order

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		tasks         int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{3, [][]int{{0, 1}, {1, 2}}}, []int{0, 1, 2}},
		{"case2", args{3, [][]int{{0, 1}, {1, 2}, {2, 0}}}, []int{}},
		{"case3", args{6, [][]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}}}, []int{0, 1, 4, 3, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.tasks, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
