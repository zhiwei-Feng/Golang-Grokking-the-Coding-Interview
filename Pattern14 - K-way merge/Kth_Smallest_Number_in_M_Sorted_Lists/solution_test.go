package kthsmallestnumberinmsortedlists

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, 5}, 4},
		{"case2", args{[][]int{{5, 8, 9}, {1, 7}}, 3}, 7},
		{"case3", args{[][]int{{1, 3}, {2, 4}}, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
