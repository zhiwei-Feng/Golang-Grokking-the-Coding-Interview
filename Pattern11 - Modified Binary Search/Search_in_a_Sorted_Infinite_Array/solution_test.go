package searchinasortedinfinitearray

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		reader ArrayReader
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{ArrayReader{[]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}}, 16}, 6},
		{"case2", args{ArrayReader{[]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}}, 11}, -1},
		{"case3", args{ArrayReader{[]int{1, 3, 8, 10, 15}}, 15}, 4},
		{"case4", args{ArrayReader{[]int{1, 3, 8, 10, 15}}, 200}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.reader, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
