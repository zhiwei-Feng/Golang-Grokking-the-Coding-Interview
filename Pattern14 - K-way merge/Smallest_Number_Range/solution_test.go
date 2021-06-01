package smallestnumberrange

import (
	"reflect"
	"testing"
)

func Test_smallestRange(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[][]int{{1, 5, 8}, {4, 12}, {7, 8, 10}}}, []int{4, 7}},
		{"case2", args{[][]int{{1, 9}, {4, 12}, {7, 10, 16}}}, []int{9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
