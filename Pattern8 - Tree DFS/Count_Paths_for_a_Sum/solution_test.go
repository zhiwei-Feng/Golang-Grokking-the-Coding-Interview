package Count_Paths_for_a_Sum

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}

	head1 := &TreeNode{10, nil, nil}
	head1.Left = &TreeNode{5, &TreeNode{3, &TreeNode{Val: 3}, &TreeNode{Val: -2}}, &TreeNode{2, nil, &TreeNode{Val: 1}}}
	head1.Right = &TreeNode{-3, nil, &TreeNode{Val: 11}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1, 8}, 3},
		{"case2", args{nil, 8}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
