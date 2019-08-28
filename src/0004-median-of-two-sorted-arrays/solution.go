package solution

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	sort.Ints(nums)
	length := len(nums)
	var ret float64
	if length%2 == 0 {
		ret = float64(nums[length/2-1]+nums[length/2]) / 2.0
	} else {
		ret = float64(nums[length/2])
	}
	return ret
}
