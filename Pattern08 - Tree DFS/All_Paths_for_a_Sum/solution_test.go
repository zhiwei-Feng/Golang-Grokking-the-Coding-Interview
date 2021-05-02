package All_Paths_for_a_Sum

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}

	head1 := &TreeNode{Val: 5}
	head1.Left = &TreeNode{Val: 4}
	head1.Left.Left = &TreeNode{11, &TreeNode{Val: 7}, &TreeNode{Val: 2}}
	head1.Right = &TreeNode{Val: 8}
	head1.Right.Left = &TreeNode{Val: 13}
	head1.Right.Right = &TreeNode{4, &TreeNode{Val: 5}, &TreeNode{Val: 1}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{head1, 22}, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
