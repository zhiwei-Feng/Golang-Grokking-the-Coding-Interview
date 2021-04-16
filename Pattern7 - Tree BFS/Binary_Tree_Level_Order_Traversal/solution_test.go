package Binary_Tree_Level_Order_Traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		head *TreeNode
	}

	var head1 *TreeNode

	head2 := &TreeNode{Val: 3}
	head2.Left = &TreeNode{Val: 9}
	head2.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{head1}, [][]int{}},
		{"case2", args{head2}, [][]int{{3}, {9, 20}, {15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
