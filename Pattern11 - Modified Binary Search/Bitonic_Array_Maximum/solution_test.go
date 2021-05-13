package bitonicarraymaximum

import "testing"

func Test_findMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 3, 8, 12, 4, 2}}, 12},
		{"case2", args{[]int{3, 8, 3, 1}}, 8},
		{"case3", args{[]int{1, 3, 8, 12}}, 12},
		{"case4", args{[]int{10, 9, 8}}, 10},
		{"case5", args{[]int{3, 8, 3}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.arr); got != tt.want {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
