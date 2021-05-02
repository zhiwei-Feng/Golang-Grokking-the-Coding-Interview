package Level_Order_Successor

import (
	"reflect"
	"testing"
)

func Test_findSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}

	head1 := &TreeNode{Val: 12}
	node7 := &TreeNode{Val: 7, Left: &TreeNode{Val: 9}}
	head1.Left = node7
	node1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 5}}
	head1.Right = node1

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1", args{head1, 12}, node7},
		{"case2", args{head1, 7}, node1},
		{"case2", args{head1, 5}, nil},
		{"case3", args{nil, 7}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSuccessor(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
