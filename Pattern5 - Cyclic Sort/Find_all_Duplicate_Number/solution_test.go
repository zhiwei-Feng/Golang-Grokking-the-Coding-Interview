package Find_all_Duplicate_Number

import (
	"reflect"
	"testing"
)

func Test_findAllDuplicateNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{3, 4, 4, 5, 5}}, []int{5, 4}},
		{"case2", args{[]int{5, 4, 7, 2, 3, 5, 3}}, []int{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllDuplicateNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllDuplicateNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
