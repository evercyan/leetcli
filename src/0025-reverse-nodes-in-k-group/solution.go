package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 开始进行交换
 *
 * 假设: head = [1, 2, 3, 4, 5, 6], k = 5
 * 初始: curr = 1, next = null, prev = [6]
 *
 * 遍历第一轮:
 * next = curr.Next, next = 2
 * curr.Next = prev, 相当于把 1 - 6 连接起来
 * prev = curr, 此时 prev = [1, 6]
 * curr = next, curr 后移, curr = 2
 * 这轮结束, prev = [1, 6]
 *
 * 故:
 * i=0, [1, 6]
 * i=1, [2, 1, 6]
 * i=2, [3, 2, 1, 6]
 * i=3, [4, 3, 2, 1, 6]
 * i=4, [5, 4, 3, 2, 1, 6]
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	count := 0
	l := head
	for count < k {
		if l == nil {
			break
		}
		l = l.Next
		count++
	}
	// 如果长度不足 k, 直接 return
	if count != k {
		return head
	}
	// 递规剩余节点
	prev := reverseKGroup(l, k)
	curr, next := head, &ListNode{}
	for i := 0; i < k; i++ {
		// 纪录 next 节点
		next = curr.Next
		// 将头节点同 prev 拼接
		curr.Next = prev
		// prev 前移
		prev = curr
		// curr 后移
		curr = next
	}
	return prev
}
