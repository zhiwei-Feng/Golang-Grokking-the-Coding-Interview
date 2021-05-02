package Sum_of_Path_Numbers

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{1, nil, nil}
	head1.Left = &TreeNode{7, nil, nil}
	head1.Right = &TreeNode{9, &TreeNode{Val: 2}, &TreeNode{Val: 9}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1}, 408},
		{"case2", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
