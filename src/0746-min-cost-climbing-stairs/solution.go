package solution

/**
 * 题解:
 * 从地面开始爬阶梯, 直到顶层, 求最小代价.
 * 将地面标号为-1, 然后逐层递增地给每一层进行编号
 * 如果cost大小为n, 则最顶层编号为n+1
 * cos[i]表示从第i-1层/第i-2层往上1/2步到第i层的代价
 *
 * 动态规划
 * 故公式可总结为
 * dp[i] = min(dp[i-2]+nums[i], dp[i-1])
 * 前 i 阶最小耗费 = min(前 i-2 阶 + 第 i-2 阶, 前 i-1 阶 + 第 i-1 阶)
 */

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[n]
}
