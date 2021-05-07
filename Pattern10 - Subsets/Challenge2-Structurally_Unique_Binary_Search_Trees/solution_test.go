package Challenge2_Structurally_Unique_Binary_Search_Trees

import (
	"fmt"
	"testing"
)

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{2}, 2},
		{"case2", args{3}, 5},
		{"case3", args{1}, 1},
		{"case4", args{4}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			//}
			got := generateTrees(tt.args.n)
			for _, v := range got {
				//print
				var queue []*TreeNode
				queue = append(queue, v)
				for len(queue) != 0 {
					o := queue[0]
					queue = queue[1:]
					fmt.Print(o.Val, " ")
					if o.Left != nil {
						queue = append(queue, o.Left)
					}
					if o.Right != nil {
						queue = append(queue, o.Right)
					}
				}
				fmt.Println()
			}
			if len(got) != tt.want {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
