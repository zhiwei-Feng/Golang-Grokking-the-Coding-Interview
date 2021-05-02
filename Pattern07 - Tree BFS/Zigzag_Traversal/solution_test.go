package Zigzag_Traversal

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{Val: 3}
	head1.Left = &TreeNode{Val: 9}
	head1.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{head1}, [][]int{{3}, {20, 9}, {15, 7}}},
		{"case2", args{nil}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
