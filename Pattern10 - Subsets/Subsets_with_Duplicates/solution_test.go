package Subsets_with_Duplicates

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[]int{1, 3, 3}}, [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}}},
		{"case2", args{[]int{1, 5, 3, 3}}, [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}, {5}, {1, 5}, {3, 5}, {1, 3, 5}, {3, 3, 5}, {1, 3, 3, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
