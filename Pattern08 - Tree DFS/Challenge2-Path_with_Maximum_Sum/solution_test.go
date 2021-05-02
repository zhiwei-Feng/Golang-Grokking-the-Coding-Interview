package Challenge2_Path_with_Maximum_Sum

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{1, nil, nil}
	head1.Left = &TreeNode{2, &TreeNode{Val: 4}, nil}
	head1.Right = &TreeNode{3, &TreeNode{Val: 5}, &TreeNode{Val: 6}}

	head2 := &TreeNode{1, nil, nil}
	head2.Left = &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	head2.Right = &TreeNode{3, nil, nil}
	head2.Right.Left = &TreeNode{5, &TreeNode{Val: 7}, &TreeNode{Val: 8}}
	head2.Right.Right = &TreeNode{6, nil, &TreeNode{Val: 9}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1}, 16},
		{"case2", args{head2}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
