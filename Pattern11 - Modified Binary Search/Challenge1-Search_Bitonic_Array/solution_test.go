package challenge1searchbitonicarray

import "testing"

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
		{"case1", args{[]int{1, 3, 8, 4, 3}, 4}, 3},
		{"case2", args{[]int{3, 8, 3, 1}, 8}, 1},
		{"case3", args{[]int{1, 3, 8, 12}, 12}, 3},
		{"case4", args{[]int{10, 9, 8}, 10}, 0},
		{"case5", args{[]int{10, 9, 8}, 7}, -1},
		{"case6", args{[]int{1, 3, 8, 12}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
