package Challenge2_Rotate_a_LinkedList

/*
Given the head of a Singly LinkedList and a number ‘k’, rotate the LinkedList to the right by ‘k’ nodes.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func rotate(head *ListNode, k int) *ListNode {
	// 2N -> O(N)
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	var (
		lastNode = head
		length   = 1
	)

	for lastNode.Next != nil {
		length++
		lastNode = lastNode.Next
	}
	lastNode.Next = head

	k = length - (k % length)
	lastNodeOfTail := head
	for i := 0; i < k-1; i++ {
		lastNodeOfTail = lastNodeOfTail.Next
	}

	head = lastNodeOfTail.Next
	lastNodeOfTail.Next = nil

	return head
}
