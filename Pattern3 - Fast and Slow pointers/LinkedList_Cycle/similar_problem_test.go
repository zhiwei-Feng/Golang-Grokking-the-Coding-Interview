package LinkedList_Cycle

import "testing"

func Test_findCycleLength(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	head1.Next.Next = &ListNode{Value: 3}
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}
	head1.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	head2.Next.Next.Next = &ListNode{Value: 4}
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head2.Next.Next.Next.Next.Next.Next = head2.Next.Next

	head3 := &ListNode{Value: 1}
	head3.Next = &ListNode{Value: 2}
	head3.Next.Next = &ListNode{Value: 3}
	head3.Next.Next.Next = &ListNode{Value: 4}
	head3.Next.Next.Next.Next = &ListNode{Value: 5}
	head3.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head3.Next.Next.Next.Next.Next.Next = head3.Next.Next.Next
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{head1}, 0},
		{"case2", args{head2}, 4},
		{"case3", args{head3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCycleLength(tt.args.head); got != tt.want {
				t.Errorf("findCycleLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
