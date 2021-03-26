package Find_all_Missing_Number

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
		{"case1", args{[]int{2, 3, 1, 8, 2, 3, 5, 1}}, []int{4, 6, 7}},
		{"case2", args{[]int{2, 4, 1, 2}}, []int{3}},
		{"case3", args{[]int{2, 3, 2, 1}}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
