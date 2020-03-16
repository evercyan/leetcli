package solution

/**
 * 目标就是判断链接是否有环, 原题中引入的 pos 实际上是自己去模拟生成一条链表带环...
 *
 * 故如 &ListNode{Val: 1, Next: &ListNode{Val: 2}}... 这样固定的用例没法单测, 需自己生成数据
 *
 * 快慢指针 & hash
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// hash
func hasCycleByHash(head *ListNode) bool {
	hash := map[int]int{}
	for head != nil {
		if _, ok := hash[head.Val]; ok {
			return true
		}
		hash[head.Val] = 1
		head = head.Next
	}
	return false
}
