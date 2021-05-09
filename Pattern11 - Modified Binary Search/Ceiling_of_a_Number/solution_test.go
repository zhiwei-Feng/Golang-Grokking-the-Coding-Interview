package Ceiling_of_a_Number

import "testing"

func Test_SearchCeilingOfANumber(t *testing.T) {
	type args struct {
		arr []int
		key int	
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 6, 10}, 6}, 1},
		{"case2", args{[]int{1, 3, 8, 10, 15}, 12}, 4},
		{"case3", args{[]int{4, 6, 10}, 17}, -1},
		{"case4", args{[]int{4, 6, 10}, -1}, 0},
	}

	for _, tt := range tests {
		if got := searchCeilingOfANumber(tt.args.arr, tt.args.key); got != tt.want {
			t.Errorf("searchCeilingOfANumber() = %v, but want %v", got, tt.want)
		}
	}
}