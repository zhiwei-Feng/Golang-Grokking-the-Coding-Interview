package Challenge1_Palindrome_LinkedList

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head1 := &ListNode{Value: 2}
	head1.Next = &ListNode{Value: 4}
	head1.Next.Next = &ListNode{Value: 6}
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 2}

	head2 := &ListNode{Value: 2}
	head2.Next = &ListNode{Value: 4}
	head2.Next.Next = &ListNode{Value: 6}
	head2.Next.Next.Next = &ListNode{Value: 4}
	head2.Next.Next.Next.Next = &ListNode{Value: 2}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 2}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{head1}, true},
		{"case2", args{head2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
