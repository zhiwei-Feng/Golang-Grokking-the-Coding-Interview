package Maximum_Sum_Subarray_of_Size_K

import "testing"

func Test_findMaxSumSubArray(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{3, []int{2, 1, 5, 1, 3, 2}}, want: 9},
		{name: "case2", args: args{2, []int{2, 3, 4, 1, 5}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSumSubArray(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("findMaxSumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
