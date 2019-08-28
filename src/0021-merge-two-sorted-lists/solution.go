package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 用 Val 比较, 小的放前, 余下一直递规
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var lret *ListNode
	if l1.Val < l2.Val {
		lret = l1
		lret.Next = mergeTwoLists(l1.Next, l2)
	} else {
		lret = l2
		lret.Next = mergeTwoLists(l1, l2.Next)
	}
	return lret
}
