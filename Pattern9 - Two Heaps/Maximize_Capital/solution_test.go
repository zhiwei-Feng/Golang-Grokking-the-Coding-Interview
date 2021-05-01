package Maximize_Capital

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{
			k:       10,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 2},
		}, 6},
		{"case2", args{
			k:       2,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 1},
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
