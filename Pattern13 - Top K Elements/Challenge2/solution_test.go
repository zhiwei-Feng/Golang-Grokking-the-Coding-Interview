package challenge2

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]byte{'a', 'a', 'a', 'b', 'c', 'c'}, 2}, 7},
		{"case2", args{[]byte{'a', 'b', 'a'}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.k); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
