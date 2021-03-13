package Challenge3_Minimum_Window_Sort

import "testing"

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 5, 3, 7, 10, 9, 12}}, 5},
		{"case2", args{[]int{1, 3, 2, 0, -1, 7, 10}}, 5},
		{"case3", args{[]int{1, 2, 3}}, 0},
		{"case4", args{[]int{3, 2, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.arr); got != tt.want {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
