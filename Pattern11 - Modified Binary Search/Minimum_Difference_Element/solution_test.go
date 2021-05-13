package minimumdifferenceelement

import "testing"

func Test_searchMinDiffElement(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 6, 10}, 7}, 6},
		{"case2", args{[]int{4, 6, 10}, 4}, 4},
		{"case3", args{[]int{1, 3, 8, 10, 15}, 12}, 10},
		{"case4", args{[]int{4, 6, 10}, 17}, 10},
		{"case5", args{[]int{4, 6, 10}, 3}, 4},
		{"case6", args{[]int{1, 3, 8, 10, 15}, 13}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMinDiffElement(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("searchMinDiffElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
