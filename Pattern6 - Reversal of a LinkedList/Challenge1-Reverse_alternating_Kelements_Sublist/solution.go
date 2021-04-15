package Challenge1_Reverse_alternating_Kelements_Sublist

/*
Given the head of a LinkedList and a number ‘k’, reverse every alternating ‘k’ sized sub-list starting from the head.
If, in the end, you are left with a sub-list with less than ‘k’ elements, reverse it too.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseAlternatingKElementsSublist(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	var (
		current  = head
		previous *ListNode
	)

	for current != nil {
		lastNodeOfSublist := current
		lastNodeOfPrev := previous

		for i := 0; i < k && current != nil; i++ {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
		}

		if lastNodeOfPrev == nil {
			head = previous
		} else {
			lastNodeOfPrev.Next = previous
		}

		lastNodeOfSublist.Next = current

		// skip k steps
		for i := 0; i < k && current != nil; i++ {
			previous = current
			current = current.Next
		}
	}

	return head
}
