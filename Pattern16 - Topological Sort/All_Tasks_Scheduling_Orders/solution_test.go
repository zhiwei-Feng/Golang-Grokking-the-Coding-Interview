package All_Tasks_Scheduling_Orders

import "testing"

func Test_printOrders(t *testing.T) {
	type args struct {
		tasks         int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{3, [][]int{{0, 1}, {1, 2}}}},
		{"case2", args{4, [][]int{{3, 2}, {3, 0}, {2, 0}, {2, 1}}}},
		{"case3", args{6, [][]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printOrders(tt.args.tasks, tt.args.prerequisites)
		})
	}
}
