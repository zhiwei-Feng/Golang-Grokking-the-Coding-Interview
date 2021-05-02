package Challenge1_Next_Interval

import (
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[][]int{{1, 3}, {2, 4}, {4, 5}}}, []int{2, 2, -1}},
		{"case2", args{[][]int{{2, 3}, {3, 4}, {5, 6}}}, []int{1, 2, -1}},
		{"case3", args{[][]int{{3, 4}, {1, 5}, {4, 6}}}, []int{2, -1, -1}},
		{"case4", args{[][]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightInterval(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
