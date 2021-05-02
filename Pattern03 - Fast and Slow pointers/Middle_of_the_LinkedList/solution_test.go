package Middle_of_the_LinkedList

import (
	"reflect"
	"testing"
)

func Test_findMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	midNode1 := &ListNode{Value: 3}
	head1.Next.Next = midNode1
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	midNode2 := &ListNode{Value: 4}
	head2.Next.Next.Next = midNode2
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head3 := &ListNode{Value: 1}
	head3.Next = &ListNode{Value: 2}
	head3.Next.Next = &ListNode{Value: 3}
	midNode3 := &ListNode{Value: 4}
	head3.Next.Next.Next = midNode3
	head3.Next.Next.Next.Next = &ListNode{Value: 5}
	head3.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head3.Next.Next.Next.Next.Next.Next = &ListNode{Value: 7}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head1}, midNode1},
		{"case2", args{head2}, midNode2},
		{"case3", args{head3}, midNode3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
