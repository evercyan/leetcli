package solution

/**
 * 题目解读
 * 求数组连续子数组和最大值
 * f(n) = max(f(n-1) + A[n], A[n]);
 */

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 先取出第一个元素值作为初始结果
	result := nums[0]
	// 纪录遍历过程中的和值
	sum := 0
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			// 如果发现之前的和值小于 0, 则重置为当前元素, 可以理解为前 i-1 项和值为负, 宁愿不加
			sum = nums[i]
		} else {
			// 加之
			sum += nums[i]
		}
		// 进行比对, 每循环一次, 拿初始结果同和值比较, 取 max
		if result < sum {
			result = sum
		}
	}
	return result
}
