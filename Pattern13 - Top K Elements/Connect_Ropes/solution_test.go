package connectropes

import "testing"

func Test_connectSticks(t *testing.T) {
	type args struct {
		sticks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 3, 11, 5}}, 33},
		{"case2", args{[]int{3, 4, 5, 6}}, 36},
		{"case3", args{[]int{1, 3, 11, 5, 2}}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectSticks(tt.args.sticks); got != tt.want {
				t.Errorf("connectSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}
