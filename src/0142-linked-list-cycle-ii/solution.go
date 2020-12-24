package solution

// 返回链表开始入环的第一个节点
// 原题中引入的 pos 实际上是自己去模拟生成一条链表带环...
// 快慢指针
// 从头节点出发, 如果链表有环, 快指针肯定可以在环内和慢指针相遇
// 没有环就不可能再相遇, 相遇必在环内

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// fast == nil 时即无环
	for fast != nil && fast.Next != nil {
		// 快慢指针开始走
		slow, fast = slow.Next, fast.Next.Next
		// 相遇时必有环, 且相遇点在环内
		if slow == fast {
			// 由快慢指针速度推导可得: 头节点到到入环点 = 相遇点到入环点
			fast = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}

	return nil
}
