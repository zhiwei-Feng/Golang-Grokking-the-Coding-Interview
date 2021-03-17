package Challenge2_Rearrange_a_LinkedList

/*
Given the head of a Singly LinkedList, write a method to modify the LinkedList
such that the nodes from the second half of the LinkedList are inserted alternately to the nodes from the first half in reverse order.
So if the LinkedList has nodes 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null, your method should return 1 -> 6 -> 2 -> 5 -> 3 -> 4 -> null.

Your algorithm should not use any extra space and the input LinkedList should be modified in-place.

Input: 2 -> 4 -> 6 -> 8 -> 10 -> 12 -> null
Output: 2 -> 12 -> 4 -> 10 -> 6 -> 8 -> null

Input: 2 -> 4 -> 6 -> 8 -> 10 -> null
Output: 2 -> 10 -> 4 -> 8 -> 6 -> null
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverse(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
		next *ListNode
	)

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reorder(head *ListNode) {
	var (
		fast, slow = head, head
		middle     *ListNode
		secondHead *ListNode
		firstHead  *ListNode
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	middle = slow
	secondHead = reverse(middle)
	firstHead = head
	for firstHead != nil && secondHead != nil {
		temp := firstHead.Next
		firstHead.Next = secondHead
		firstHead = temp

		temp = secondHead.Next
		secondHead.Next = firstHead
		secondHead = temp
	}

	// 注意结尾的处理
	if firstHead != nil {
		firstHead.Next = nil
	}
}
