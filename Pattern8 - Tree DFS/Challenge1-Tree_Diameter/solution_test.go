package Challenge1_Tree_Diameter

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{1, nil, nil}
	head1.Left = &TreeNode{2, &TreeNode{Val: 4}, nil}
	head1.Right = &TreeNode{3, &TreeNode{Val: 5}, &TreeNode{Val: 6}}

	head2 := &TreeNode{1, &TreeNode{Val: 2}, nil}
	head2.Right = &TreeNode{3, nil, nil}
	head2.Right.Left = &TreeNode{5, &TreeNode{Val: 7}, &TreeNode{8, &TreeNode{Val: 10}, nil}}
	head2.Right.Right = &TreeNode{6, nil, &TreeNode{9, nil, &TreeNode{Val: 11}}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1}, 5},
		{"case2", args{head2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
