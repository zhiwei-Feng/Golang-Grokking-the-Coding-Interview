package Challenge2_Rearrange_a_LinkedList

import (
	"testing"
)

func Test_reorder(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := &ListNode{Value: 2}
	head1.Next = &ListNode{Value: 4}
	head1.Next.Next = &ListNode{Value: 6}
	head1.Next.Next.Next = &ListNode{Value: 8}
	head1.Next.Next.Next.Next = &ListNode{Value: 10}
	head1.Next.Next.Next.Next.Next = &ListNode{Value: 12}

	head2 := &ListNode{Value: 2}
	head2.Next = &ListNode{Value: 4}
	head2.Next.Next = &ListNode{Value: 6}
	head2.Next.Next.Next = &ListNode{Value: 8}
	head2.Next.Next.Next.Next = &ListNode{Value: 10}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{head1}, []int{2, 12, 4, 10, 6, 8}},
		{"case2", args{head2}, []int{2, 10, 4, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorder(tt.args.head)
			travel := tt.args.head
			tempArray := make([]int, 0, len(tt.want))
			for travel != nil {
				tempArray = append(tempArray, travel.Value)
				travel = travel.Next
			}
			resultFlag := true
			i := 0
			for ; i < len(tt.want); i++ {
				if tt.want[i] != tempArray[i] {
					resultFlag = false
					break
				}
			}
			if i != len(tt.want) {
				resultFlag = false
			}
			if !resultFlag {
				t.Errorf("reorder() = %v, want %v", tempArray, tt.want)
			}
		})
	}
}
