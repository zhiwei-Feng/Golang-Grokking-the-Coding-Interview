package Challenge1_Reverse_alternating_Kelements_Sublist

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseAlternatingKElementsSublist(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	head1 := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	head1.Next = node2
	head1.Next.Next = &ListNode{Value: 3}
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}
	head1.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head1.Next.Next.Next.Next.Next.Next = &ListNode{Value: 7}
	head1.Next.Next.Next.Next.Next.Next.Next = &ListNode{Value: 8}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	head2.Next.Next.Next = &ListNode{Value: 4}
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head2.Next.Next.Next.Next.Next.Next = &ListNode{Value: 7}
	head2.Next.Next.Next.Next.Next.Next.Next = &ListNode{Value: 8}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head1, 2}, node2},
		{"case2", args{head2, 1}, head2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseAlternatingKElementsSublist(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseAlternatingKElementsSublist() = %v, want %v", got, tt.want)
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
