package Reverse_a_Sublist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	/**
	current, prev, next用于反转子链表
	tail 记录子链表初始时的头结点，用于反转后连接右侧剩余节点
	*/
	var (
		current, prev, next, tail *ListNode
		LeftNode                  *ListNode
		RightNode                 *ListNode
		p                         = head
		i                         = 1
	)

	// LeftNode表示子链表的前一个节点
	for i < left {
		LeftNode = p
		p = p.Next
		i++
	}

	current = p
	tail = current

	for i < right {
		p = p.Next
		i++
	}

	RightNode = p.Next
	// 断开子链表和右侧剩余节点的连接，方便子链表反转
	p.Next = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	// 连接上右侧剩余节点
	tail.Next = RightNode

	if LeftNode != nil {
		// 说明此时原链表的头结点并未反转
		LeftNode.Next = prev
		return head
	} else {
		// 头结点被反转
		return prev
	}

}
