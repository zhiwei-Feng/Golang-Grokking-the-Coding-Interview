package subsetsum

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		num []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 2, 3, 7}, 6}, true},
		{"case2", args{[]int{1, 2, 7, 1, 5}, 10}, true},
		{"case3", args{[]int{1, 3, 4, 8}, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.num, tt.args.sum); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
