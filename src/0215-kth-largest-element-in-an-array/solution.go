package solution

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {

	// 倒序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i, num := range nums {
		if i+1 == k {
			return num
		}
	}

	return -1
}
