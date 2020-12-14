package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	// 此处不可 prev := &ListNode{}, 会有零值
	var prev *ListNode
	for head != nil {
		// 记录当前节点的下一个节点
		next := head.Next
		// 将当前节点指向 prev
		head.Next = prev
		// prev 和 head 都前进一位
		prev = head
		head = next
	}

	return prev
}

// 递规
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	// curr 就是最后一个节点
//	curr := reverseList(head.Next)
//
//	// 如果链表是 1->2->3->4->5, 那么此时的 curr 就是 5
//	// 而 head 是 4, head.Next 是 5, head.Next.Next 是 nil
//	// head.Next.Next = head 则是 5->4
//	head.Next.Next = head
//	// 防止链表循环, 需要将 head.Next 设置为 nil
//	head.Next = nil
//
//	return curr
//}
