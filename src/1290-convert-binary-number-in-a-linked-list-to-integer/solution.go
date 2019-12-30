package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var nums []int
	nums = append(nums, head.Val)
	for head.Next != nil {
		// 从头部插入
		nums = append([]int{head.Next.Val}, nums...)
		head = head.Next
	}
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		// 2^i 次方
		num += 1 << i
	}
	return num
}
