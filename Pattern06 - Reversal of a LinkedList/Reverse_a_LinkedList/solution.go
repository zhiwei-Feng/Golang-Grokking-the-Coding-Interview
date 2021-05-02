package Reverse_a_LinkedList

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverse(head *ListNode) *ListNode {
	var (
		current  = head
		previous *ListNode
		next     *ListNode
	)

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}
