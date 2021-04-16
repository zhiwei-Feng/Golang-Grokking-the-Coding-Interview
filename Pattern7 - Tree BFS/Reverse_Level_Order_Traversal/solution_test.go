package Reverse_Level_Order_Traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
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
		{"case2", args{head2}, [][]int{{15, 7}, {9, 20}, {3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
