package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 链表转数组, 双指针比对
// 2. 翻转链表, 新老链表比对

func isPalindrome(head *ListNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	l := len(nums)
	for i := 0; i < l/2; i++ {
		if nums[i] != nums[l-1-i] {
			return false
		}
	}

	return true
}
