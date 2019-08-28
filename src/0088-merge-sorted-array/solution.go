package solution

/**
 * 题目解读
 * 归并排序中的 merge 操作, 只是索引是从后面开始进行归并
 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 最终数组最后一个元素下标
	tail := m + n - 1
	// nums1, nums2 最后一个元素下标
	m -= 1
	n -= 1
	// 遍历, 直至有一个数组消耗完毕
	for m >= 0 && n >= 0 {
		// 哪个大, 放至尾部, tail--
		if nums1[m] >= nums2[n] {
			nums1[tail] = nums1[m]
			m--
		} else {
			nums1[tail] = nums2[n]
			n--
		}
		tail--
	}
	// 如果是 m > 0, 则 n 被消耗完, m 前置元素已在 m 中, 无需要处理
	// 如果是 n > 0, 则 m 被消耗完, 需要将 n 前置元素放置 m 中
	for n >= 0 {
		nums1[n] = nums2[n]
		n--
	}
	return
}
