package sumofelements

import "testing"

func Test_findSumOfElements(t *testing.T) {
	type args struct {
		nums []int
		k1   int
		k2   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 3, 12, 5, 15, 11}, 3, 6}, 23},
		{"case2", args{[]int{3, 5, 8, 7}, 1, 4}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumOfElements(tt.args.nums, tt.args.k1, tt.args.k2); got != tt.want {
				t.Errorf("findSumOfElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
