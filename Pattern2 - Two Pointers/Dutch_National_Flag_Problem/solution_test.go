package Dutch_National_Flag_Problem

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 0, 2, 1, 0}}, []int{0, 0, 1, 1, 2}},
		{"case2", args{[]int{2, 2, 0, 1, 2, 0}}, []int{0, 0, 1, 2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("sort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
