package Challenge3_Find_the_First_K_Missing_Positive_Numbers

import (
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{3, -1, 4, 5, 5}, 3}, []int{1, 2, 6}},
		{"case2", args{[]int{2, 3, 4}, 3}, []int{1, 5, 6}},
		{"case3", args{[]int{-2, -3, 4}, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
