package solution

// 关注 0 值, 遇 0 时看前置位能跳过 0
// 令 dp[i] 表示第 i 个元素是否可跳
// dp[i-1] 表示 i-1 个元素是否可跳, nums[i-1] >= (i - (i-1)) 表示如果 i-1 个元素可跳的话, 该元素值必须 >= 位置差, 也就是前个要 >= 1, 前 2 个要 >= 2
// 最终 dp[i] 取 dp[0] - dp[i-1] 的或值, 有一可到即可
// 有 dp[i] = ( dp[i-1] && nums[i-1] >= (i - (i-1)) ) || ( dp[i-2] && nums[i-2] >= (i - (i-2)) ) ...

func canJump(nums []int) bool {
	l := len(nums)

	dp := make([]bool, l)
	dp[0] = true

	for i := 1; i < l; i++ {
		dp[i] = false
		for j := 0; j < i; j++ {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}

	return dp[l-1]
}
