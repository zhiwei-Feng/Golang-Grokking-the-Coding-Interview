package Challenge2_Maximum_CPU_Load

import "testing"

func Test_findMaxCPULoad(t *testing.T) {
	type args struct {
		jobs []Job
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]Job{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}}, 7},
		{"case2", args{[]Job{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}}, 15},
		{"case3", args{[]Job{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}}, 8},
		{"case4", args{[]Job{}}, 0},
		{"case5", args{[]Job{{1, 4, 3}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxCPULoad(tt.args.jobs); got != tt.want {
				t.Errorf("findMaxCPULoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
