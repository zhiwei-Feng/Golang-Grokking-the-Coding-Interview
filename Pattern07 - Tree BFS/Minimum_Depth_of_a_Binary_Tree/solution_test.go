package Minimum_Depth_of_a_Binary_Tree

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{Val: 3}
	head1.Left = &TreeNode{Val: 9}
	head1.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	head2 := &TreeNode{Val: 1}
	head2.Left = &TreeNode{Val: 2}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1}, 2},
		{"case2", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
