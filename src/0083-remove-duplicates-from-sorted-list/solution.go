package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	for {
		if tmp != nil && tmp.Next != nil {
			if tmp.Val == tmp.Next.Val {
				tmp.Next = tmp.Next.Next
			} else {
				tmp = tmp.Next
			}
		} else {
			break
		}
	}
	return head
}
