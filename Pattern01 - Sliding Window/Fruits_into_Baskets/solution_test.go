package Fruits_into_Baskets

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{arr: []byte{'A', 'B', 'C', 'A', 'C'}}, 3},
		{"case2", args{arr: []byte{'A', 'B', 'C', 'B', 'B', 'C'}}, 5},
		{"case3", args{arr: []byte{'A', 'R', 'A', 'A', 'C', 'I'}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.arr); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
