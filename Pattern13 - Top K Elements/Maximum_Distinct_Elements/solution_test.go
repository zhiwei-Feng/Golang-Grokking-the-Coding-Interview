package maximumdistinctelements

import "testing"

func Test_findMaximumDistinctElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{7, 3, 5, 8, 5, 3, 3}, 2}, 3},
		{"case2", args{[]int{3, 5, 12, 11, 12}, 3}, 2},
		{"case3", args{[]int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 5}, 2}, 3},
		{"case4", args{[]int{12, 3, 4, 4, 4}, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumDistinctElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaximumDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
