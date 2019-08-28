package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 注意题目: 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 下一节点
	next := head.Next
	// 跳过 next, 拼接 next.Next 节点
	head.Next = swapPairs(next.Next)
	// 完成节点交换
	next.Next = head
	return next
}

/*
func swapPairs(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil {
		if tmp.Next == nil {
			break
		}
		tmp.Val, tmp.Next.Val = tmp.Next.Val, tmp.Val
		if tmp.Next.Next == nil {
			break
		}
		tmp = tmp.Next.Next
	}
	return head
}
*/
