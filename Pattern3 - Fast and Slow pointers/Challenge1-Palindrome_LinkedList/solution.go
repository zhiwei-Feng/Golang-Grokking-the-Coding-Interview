package Challenge1_Palindrome_LinkedList

/*
Given the head of a Singly LinkedList, write a method to check if the LinkedList is a palindrome or not.

Your algorithm should use constant space and the input LinkedList should be in the original form once the algorithm is finished.
The algorithm should have O(N) time complexity where ‘N’ is the number of nodes in the LinkedList.

Input: 2 -> 4 -> 6 -> 4 -> 2 -> null
Output: true

Input: 2 -> 4 -> 6 -> 4 -> 2 -> 2 -> null
Output: false
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverse(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
	)

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	// space:O(1), time:O(n), 链表节点顺序保持不变
	var (
		fast           = head
		slow           = head
		middle         *ListNode
		secondListHead *ListNode
		firstListHead  *ListNode
		copySecond     *ListNode
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	middle = slow

	secondListHead = reverse(middle)
	copySecond = secondListHead
	firstListHead = head
	for firstListHead != nil && secondListHead != nil {
		if firstListHead.Value != secondListHead.Value {
			break
		}
		firstListHead = firstListHead.Next
		secondListHead = secondListHead.Next
	}

	reverse(copySecond)

	// 通过打印链表可以看到链表还是保持初始状态
	//for head != nil {
	//	fmt.Print(head.Value, " ")
	//	head = head.Next
	//}
	if firstListHead == nil || secondListHead == nil {
		return true
	}
	return false
}
