package Connect_Level_Order_Siblings

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}

	head1 := &Node{Val: 1}
	head1.Left = &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}
	head1.Right = &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"case1", args{head1}, head1},
		{"case2", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
