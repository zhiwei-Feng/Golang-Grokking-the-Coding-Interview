package Start_of_LinkedList_Cycle

import (
	"reflect"
	"testing"
)

func Test_findCycleStart(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	cycleStart1 := &ListNode{Value: 3}
	head1.Next.Next = cycleStart1
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}
	head1.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	cycleStart2 := &ListNode{Value: 4}
	head2.Next.Next.Next = cycleStart2
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head2.Next.Next.Next.Next.Next.Next = head2.Next.Next.Next

	head3 := &ListNode{Value: 1}
	head3.Next = &ListNode{Value: 2}
	head3.Next.Next = &ListNode{Value: 3}
	head3.Next.Next.Next = &ListNode{Value: 4}
	head3.Next.Next.Next.Next = &ListNode{Value: 5}
	head3.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head3.Next.Next.Next.Next.Next.Next = head3

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{head1}, cycleStart1},
		{"case2", args{head2}, cycleStart2},
		{"case3", args{head3}, head3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCycleStart(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCycleStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
