package Challenge1

import "testing"

func Test_sequenceReconstruction(t *testing.T) {
	type args struct {
		org  []int
		seqs [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}}, false},
		{"case2", args{[]int{1, 2, 3}, [][]int{{1, 2}}}, false},
		{"case3", args{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}}}, true},
		{"case4", args{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {2, 4}}}, false},
		{"case5", args{[]int{3, 1, 4, 2, 5}, [][]int{{3, 1, 5}, {1, 4, 2, 5}}}, true},
		{"case6", args{[]int{}, [][]int{{3, 1, 5}, {1, 4, 2, 5}}}, false},
		{"case7", args{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {1, 3}, {2, 3}, {4, 2}}}, false},
		{"case8", args{[]int{1, 2, 3, 4}, [][]int{{1, 3}, {3, 4}, {4, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequenceReconstruction(tt.args.org, tt.args.seqs); got != tt.want {
				t.Errorf("sequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
