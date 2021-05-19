package challenge1

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		image [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[][]int{}}, [][]int{}},
		{"case2", args{[][]int{{1, 0, 1}, {1, 1, 1}, {0, 1, 1}}}, [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}},
		{"case3", args{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
