package solution

// 动态规划
// dp[i, j] 表示从 (0, 0) 到 (i, j) 的路径数量, 每次只能向下或者向右
// 所以有 dp[i, j] = dp[i-1, j] + dp[i, j-1], 最终取 dp[m-1, n-1]
// 注意, 第 1 行和第 1 列, 只有一条路径, 需初始化为 1

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
