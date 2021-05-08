package orderagnosticbinarysearch

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		arr []int
		key int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 6, 10}, 10}, 2},
		{"case2", args{[]int{1, 2, 3, 4, 5, 6, 7}, 5}, 4},
		{"case3", args{[]int{10, 6, 4}, 10}, 0},
		{"case4", args{[]int{10, 6, 4}, 4}, 2},
		{"case5", args{[]int{4}, 4}, 0},
		{"case6", args{[]int{10, 6, 4}, 3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.arr, tt.args.key)
			if got != tt.want {
				t.Errorf("search() = %v, want %v\n", got, tt.want)
			}
		})
	}
}
