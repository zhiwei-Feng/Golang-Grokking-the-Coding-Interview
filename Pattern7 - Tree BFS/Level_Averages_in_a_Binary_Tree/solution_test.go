package Level_Averages_in_a_Binary_Tree

import (
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	head1 := &TreeNode{Val: 3}
	head1.Left = &TreeNode{Val: 9}
	head1.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"case1", args{head1}, []float64{3, 14.5, 11}},
		{"case2", args{nil}, []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
