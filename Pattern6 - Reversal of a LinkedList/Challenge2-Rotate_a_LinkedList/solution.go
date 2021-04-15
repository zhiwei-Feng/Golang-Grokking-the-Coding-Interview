package Challenge2_Rotate_a_LinkedList

/*
Given the head of a Singly LinkedList and a number ‘k’, rotate the LinkedList to the right by ‘k’ nodes.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func rotate(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	var (
		tail   *ListNode
		length = 0
	)

	move := head
	for move != nil {
		tail = move
		move = move.Next
		length++
	}

	k = length - (k % length)

	move = head
	for i := 0; i < k; i++ {
		move = move.Next
	}
	tail.Next = head
	lastNodeOfTail := head
	for i := 0; i < k-1; i++ {
		lastNodeOfTail = lastNodeOfTail.Next
	}

	head = lastNodeOfTail.Next
	lastNodeOfTail.Next = nil

	return head
}
