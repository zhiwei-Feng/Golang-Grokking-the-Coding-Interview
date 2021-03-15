package Start_of_LinkedList_Cycle

/*
Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle.
*/

type ListNode struct {
	Value int
	Next  *ListNode
}

func findCycleStart(head *ListNode) *ListNode {
	// version 1 通过记录走过的位置, 这种方法有比较大的空间消耗
	//moveRecord := make(map[*ListNode]bool)
	//move := head
	//for {
	//	if _, ok := moveRecord[move]; ok {
	//		break
	//	}
	//	moveRecord[move] = true
	//	move = move.Next
	//}
	//return move

	// version 2 利用cycle的长度来寻找
	// step1 获取cycle的长度k
	var (
		fast, slow = head, head
		p1, p2     = head, head
		cycleLen   = 0
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			move := fast
			move = move.Next
			cycleLen++
			for move != fast {
				move = move.Next
				cycleLen++
			}
			break
		}
	}

	// step2 将p2向前移动k个位置
	for i := 0; i < cycleLen; i++ {
		p2 = p2.Next
	}

	// step3 同步移动p1,p2，相遇时即cycle的start node
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
