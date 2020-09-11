package solution

/**
 * 动态规划方程
 *
 * dp[i] = max(dp[i-2]+nums[i], dp[i-1])
 * 前 i 间最大收益 = max(前 i-2 收益 + 第 i 间收益, 前 i-1 间收益)
 */
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// 返回数组中的最大值
func maxList(l []int) int {
	max := 0
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

/*
func robDP1(nums []int) int {
	// 没有房间
	if len(nums) == 0 {
		return 0
	}
	// 只有 1 间
	if len(nums) == 1 {
		return nums[0]
	}
	// 只有 2 间, 无法比对 dp[i-2], 直接 max 判断返回
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	// 前 1 间最大收益
	dp[0] = nums[0]
	// 前 2 间最大收益
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// 前 i 间的最大收益
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
*/

/**
 * 动态规划优化, 整个循环过程只出现 i-2, i-1, i 三个变量
 */
func rob(nums []int) int {
	if len(nums) <= 2 {
		return maxList(nums)
	}
	// 前 i-2, i-1, n 间的最大收益(初始状态)
	i2, i1, ret := nums[0], max(nums[0], nums[1]), 0
	for i := 2; i < len(nums); i++ {
		ret = max(i2+nums[i], i1)
		// 变量交换偏移
		i2 = i1
		i1 = ret
	}
	return ret
}
