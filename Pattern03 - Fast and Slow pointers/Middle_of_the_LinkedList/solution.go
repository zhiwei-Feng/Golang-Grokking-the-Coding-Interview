package Middle_of_the_LinkedList

/*
Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.
If the total number of nodes in the LinkedList is even, return the second middle node.

Input: 1 -> 2 -> 3 -> 4 -> 5 -> null
Output: 3

Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
Output: 4

Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> null
Output: 4
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func findMiddle(head *ListNode) *ListNode {
	var (
		fast = head
		slow = head
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
