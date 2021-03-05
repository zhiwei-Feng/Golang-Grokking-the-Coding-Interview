package Pair_with_Target_Sum

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		arr       []int
		targetSum int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 6}, 6}, 1, 3},
		{"case2", args{[]int{2, 5, 9, 11}, 11}, 0, 2},
		{"case3", args{[]int{2, 5, 9, 11}, 10}, -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search(tt.args.arr, tt.args.targetSum)
			if got != tt.want {
				t.Errorf("search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
