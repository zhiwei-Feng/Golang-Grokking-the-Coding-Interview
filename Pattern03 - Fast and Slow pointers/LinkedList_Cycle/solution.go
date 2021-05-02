package LinkedList_Cycle

/*
Given the head of a Singly LinkedList, write a function to determine if the LinkedList has a cycle in it or not.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func hasCycle(head *ListNode) bool {
	// 同时出发
	var (
		slow = head
		fast = head
	)

	for fast != nil && fast.Next != nil {
		// fast走两步，slow走一步，这样两者的距离每次缩小1个位置（有环）
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}

	return false
}
