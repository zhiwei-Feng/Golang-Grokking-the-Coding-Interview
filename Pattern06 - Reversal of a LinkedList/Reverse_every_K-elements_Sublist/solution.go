package Reverse_every_K_elements_Sublist

/*
Given the head of a LinkedList and a number ‘k’, reverse every ‘k’ sized sub-list starting from the head.
If, in the end, you are left with a sub-list with less than ‘k’ elements, reverse it too.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverse(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	var (
		current  = head
		previous *ListNode
	)

	for {
		lastNodeOfPrevPart := previous
		lastNodeOfSublist := current
		var next *ListNode
		for i := 0; i < k && current != nil; i++ {
			next = current.Next
			current.Next = previous
			previous = current
			current = next
		}
		if lastNodeOfPrevPart != nil {
			lastNodeOfPrevPart.Next = previous
		} else {
			head = previous
		}

		lastNodeOfSublist.Next = current
		if current == nil {
			break
		}
		previous = lastNodeOfSublist
	}

	return head
}
