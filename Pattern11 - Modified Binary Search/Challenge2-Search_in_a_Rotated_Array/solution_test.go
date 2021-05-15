package challenge2searchinarotatedarray

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"case2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"case3", args{[]int{1}, 0}, -1},
		{"case4", args{[]int{10, 15, 1, 3, 8}, 15}, 1},
		{"case5", args{[]int{4, 5, 7, 9, 10, -1, 2}, 10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
