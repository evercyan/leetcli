package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	var result *ListNode
	in := 0
	for p1 != nil && p2 != nil {
		p1.Val += p2.Val + in
		in = 0
		if p1.Val > 9 {
			p1.Val %= 10
			in = 1
		}
		result = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	if p2 != nil {
		result.Next = p2
		p1 = p2
	}
	for in > 0 {
		if p1 == nil {
			p1 = &ListNode{
				Val: 0,
			}
			result.Next = p1
		}
		p1.Val += in
		in = 0
		if p1.Val > 9 {
			p1.Val %= 10
			in = 1
		}
		result = p1
		p1 = p1.Next
	}
	return l1
}
