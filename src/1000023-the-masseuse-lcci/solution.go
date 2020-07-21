package solution

/**
 * 动态规则
 *
 * dp[i] 表示前几号预约最大时长, 满足
 * dp[i] = max(dp[i-1], dp[i-2] + nums[i])
 */

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func massage(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	// 前 1 号预约最大时间
	dp[0] = nums[0]
	// 前 2 号预约最大时间
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[l-1]
}

/**
 * 动态规划优化, 整个循环过程只出现 i-2, i-1, i 三个变量
 */

func maxList(l []int) int {
	max := 0
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

func massage2(nums []int) int {
	if len(nums) <= 2 {
		return maxList(nums)
	}
	// 前 i-2, i-1, n 间的最大收益(初始状态)
	i_2, i_1, result := nums[0], max(nums[0], nums[1]), 0

	for i := 2; i < len(nums); i++ {
		result = max(i_2+nums[i], i_1)
		// 变量交换偏移
		i_2 = i_1
		i_1 = result
	}

	return result
}
