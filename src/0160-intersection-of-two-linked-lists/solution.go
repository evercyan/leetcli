package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针遍历
// 4 1 8 4 5   5 0 1 8 4 5
// 5 0 1 8 4 5   4 1 8 4 5

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB

	// 注意此处是 a != b, 而非 a.Val != b.Val, 看地址, 而非 Val
	for a != b {
		// 遍历到尾部后, 继续遍历, 抵消长度差
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
