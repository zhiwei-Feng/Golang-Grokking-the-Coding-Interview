package Binary_Tree_Path_Sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}

	head1 := &TreeNode{12, nil, nil}
	head1.Left = &TreeNode{7, &TreeNode{Val: 9}, nil}
	head1.Right = &TreeNode{1, &TreeNode{Val: 10}, &TreeNode{Val: 5}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{nil, 1}, false},
		{"case2", args{head1, 23}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
