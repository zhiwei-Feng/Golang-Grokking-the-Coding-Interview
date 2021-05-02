package Find_the_Missing_Number

import "testing"

func Test_findMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 0, 3, 1}}, 2},
		{"case2", args{[]int{8, 3, 5, 2, 4, 6, 0, 1}}, 7},
		{"case3", args{[]int{1, 3, 2, 0}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("findMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
