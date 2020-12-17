package solution

/**
 * 题目分类: 分治算法, 堆
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	llen := len(lists)
	if llen == 0 {
		return nil
	}
	if llen == 1 {
		return lists[0]
	}
	if llen == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	var l, r []*ListNode
	for i := 0; i < llen; i++ {
		if i <= llen/2 {
			l = append(l, lists[i])
		} else {
			r = append(r, lists[i])
		}
	}
	return mergeTwoLists(mergeKLists(l), mergeKLists(r))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{}
	if l1.Val < l2.Val {
		l = l1
		l.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l = l2
		l.Next = mergeTwoLists(l1, l2.Next)
	}
	return l
}
