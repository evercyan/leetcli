package solution

// 暴力解法, 生成 hash, 取多数者
// 但此题明确给定数组一定会有多数元素占总数一半以上, 故

func majorityElement(nums []int) int {
	// major 多数元素, count 出现次数
	major, count := 0, 0

	for _, num := range nums {
		// 如果次数为 0, 被抵消完, 换元素
		// 总归多数元素占一半以上, 会留存下来
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}
