package Path_With_Given_Sequence

import "testing"

func Test_isValidSequence(t *testing.T) {
	type args struct {
		root *TreeNode
		arr  []int
	}

	head1 := &TreeNode{1, nil, nil}
	head1.Left = &TreeNode{7, nil, nil}
	head1.Right = &TreeNode{9, &TreeNode{Val: 2}, &TreeNode{Val: 9}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{head1, []int{1, 9, 9}}, true},
		{"case2", args{head1, []int{1, 9, 9, 0}}, false},
		{"case3", args{nil, []int{1, 9, 9}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSequence(tt.args.root, tt.args.arr); got != tt.want {
				t.Errorf("isValidSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
