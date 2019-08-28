package solution

/**
 * 题目要求是只遍历一遍, 故令 l1, l2 := head, head
 * 让 l1 先走 n, 然后 l1, l2, 同时走, l1 到头后, l2 的点就是倒数 n 的点
 *
 * Attention: 注意 golang 的指针复制
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l1, l2 := head, head
	for ; n > 0; n-- {
		l1 = l1.Next
	}
	if l1 == nil {
		head = head.Next
	} else {
		for l1.Next != nil {
			l1 = l1.Next
			l2 = l2.Next
		}
		l2.Next = l2.Next.Next
	}
	return head
}
