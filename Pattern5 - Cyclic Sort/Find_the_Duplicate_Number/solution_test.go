package Find_the_Duplicate_Number

import "testing"

func Test_findDuplicateNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 4, 4, 3, 2}}, 4},
		{"case2", args{[]int{2, 1, 3, 3, 5, 4}}, 3},
		{"case3", args{[]int{2, 4, 1, 4, 4}}, 4},
		//{"case4", args{[]int{2, 4, 3, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateNumber(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicateNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
