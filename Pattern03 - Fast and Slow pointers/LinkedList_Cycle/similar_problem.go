package LinkedList_Cycle

/*
Given the head of a LinkedList with a cycle, find the length of the cycle.
*/

func findCycleLength(head *ListNode) int {
	var (
		slow = head
		fast = head
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			// 计算长度
			return calculateCircleLength(slow)
		}
	}

	return 0
}

func calculateCircleLength(slow *ListNode) int {
	// slow in a circle
	var (
		current = slow
		length  = 0
	)
	// first step
	current = current.Next
	length++
	for current != slow {
		current = current.Next
		length++
	}

	return length
}
