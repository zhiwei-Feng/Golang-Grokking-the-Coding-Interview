package Challenge1_Find_the_Corrupt_Pair

import (
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{3, 1, 2, 5, 2}}, []int{2, 4}},
		{"case2", args{[]int{3, 1, 2, 3, 6, 4}}, []int{3, 5}},
		{"case3", args{[]int{3, 1, 2, 4}}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
