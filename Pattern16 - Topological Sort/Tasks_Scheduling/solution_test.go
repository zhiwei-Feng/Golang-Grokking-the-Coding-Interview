package Tasks_Scheduling

import "testing"

func Test_isSchedulingPossible(t *testing.T) {
	type args struct {
		tasks         int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{3, [][]int{{0, 1}, {1, 2}}}, true},
		{"case2", args{3, [][]int{{0, 1}, {1, 2}, {2, 0}}}, false},
		{"case3", args{6, [][]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSchedulingPossible(tt.args.tasks, tt.args.prerequisites); got != tt.want {
				t.Errorf("isSchedulingPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
