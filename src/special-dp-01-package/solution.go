package solution

/**
 * dp[i][j] 表示前 i 件物品放入容量为 j 时可获得的最大价值
 *
 * 1. i = 0 时, 表示前 0 个物品, 也就是没有物品, 其最大价值为 0, 即初始化: dp[0][0...c] = 0
 * 2. i > 0 时, dp[i][j] 存在两种情况:
 * 2.1 j >= w[i], 也就是能装入的情况下: dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]] + v[i])
 */

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func dpPackage(cap int, weight []int, value []int) int {
	num := len(weight)
	if num == 0 {
		return 0
	}

	dp := make([][]int, num+1)
	for i := 0; i <= num; i++ {
		dp[i] = make([]int, cap+1)
	}

	for i := 0; i <= cap; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= num; i++ {
		for j := 1; j <= cap; j++ {
			if j >= weight[i-1] {
				// 放得下时, 存入第 i 件商品
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i-1])
			} else {
				// 不存入第 i 件商品
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[num][cap]
}
