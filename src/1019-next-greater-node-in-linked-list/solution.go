package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 题解
 * 即下一个最大元素的链表版本
 * 下一个最大元素, 即每个元素右边第一个比它大的元素, 如果没有输出 -1
 * input : [1, 3, 2, 4, 5]
 * answer: [3, 4, 4, 5, -1]
 */

func nextLargerNodes(head *ListNode) []int {
	// 先将链表转为数组
	list, ret := []int{}, []int{}
	for {
		if head == nil {
			break
		}
		list = append(list, head.Val)
		head = head.Next
	}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			// 右边遇到第一个比它大的, 就认定它, 跟它走吧~
			if list[j] > list[i] {
				ret = append(ret, list[j])
				break
			}
			// 循环至最右, 仍无比当前元素大的
			if j == len(list)-1 {
				ret = append(ret, 0)
			}
		}
	}
	// 最右元素很孤独, 右边没比它大的元素了...
	ret = append(ret, 0)
	return ret
}
