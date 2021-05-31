package kthsmallestnumberinasortedmatrix

import "testing"

func Test_kthSmallestByBinarySearch(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8}, 13},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestByBinarySearch(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestByBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
