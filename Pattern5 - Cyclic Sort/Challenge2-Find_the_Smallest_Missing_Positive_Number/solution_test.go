package Challenge2_Find_the_Smallest_Missing_Positive_Number

import "testing"

func Test_findNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{-3, 1, 5, 4, 2}}, 3},
		{"case2", args{[]int{3, -2, 0, 1, 2}}, 4},
		{"case3", args{[]int{3, 2, 5, 1}}, 4},
		{"case4", args{[]int{3, 2, 4, 1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumber(tt.args.nums); got != tt.want {
				t.Errorf("findNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
