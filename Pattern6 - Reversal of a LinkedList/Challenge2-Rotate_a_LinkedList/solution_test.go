package Challenge2_Rotate_a_LinkedList

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}
	head1.Next.Next = node3
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	node4 := &ListNode{Value: 4}
	head2.Next.Next.Next = node4
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head2, 3}, node4},
		{"case2", args{head1, 8}, node3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			} else {
				p := got
				for p != nil {
					fmt.Printf("%d ", p.Value)
					p = p.Next
				}
				fmt.Println()
			}
		})
	}
}
