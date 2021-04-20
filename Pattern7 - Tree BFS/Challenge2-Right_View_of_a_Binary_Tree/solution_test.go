package Challenge2_Right_View_of_a_Binary_Tree

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{12, nil, nil}
	head1.Left = &TreeNode{7, &TreeNode{9, &TreeNode{Val: 3}, nil}, nil}
	head1.Right = &TreeNode{1, &TreeNode{Val: 10}, &TreeNode{Val: 5}}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{nil}, []int{}},
		{"case2", args{head1}, []int{12, 1, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
