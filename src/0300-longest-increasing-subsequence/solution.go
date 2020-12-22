package solution

// 动态规划
// dp[i] 表示下标为 i 的元素和前置位元素的最大长度子序列
// 如果 nums[i] > nums[i-1], 即有 dp[i] = dp[i-1]+1
// 即, 满足 nums[i] > nums[i-1] 时, dp[i] = max(dp[i-1]+1, dp[i-2]+1, ....)
// 最终, 取 dp 中的最大值即可

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 1
	max := dp[0]
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
