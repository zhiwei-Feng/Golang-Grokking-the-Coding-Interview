package Challenge3_Cycle_in_a_Circular_Array

import "testing"

func Test_loopExists(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 2, -1, 2, 2}}, true},
		{"case2", args{[]int{2, 2, -1, 2}}, true},
		{"case3", args{[]int{2, 1, -1, -2}}, false},
		{"case4", args{[]int{2, 1, 4, -2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loopExists(tt.args.arr); got != tt.want {
				t.Errorf("loopExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
