package solution

// 典型动态规划
// dp[i, j] 表示左上角到其点的路径总和, 因为每次只能向下或者向右移动一步, 故有
// dp[i, j] = min(dp[i-1, j](上点), dp[i, j-1](左点)) + grid[i][j]

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prev := 0
			if i == 0 && j > 0 {
				// 是第一行, 非第一列, 即有 左边节点
				prev = dp[i][j-1]
			} else if i > 0 && j == 0 {
				// 非第一行, 是第一列, 即有 上边节点
				prev = dp[i-1][j]
			} else if i > 0 && j > 0 {
				// 非第一行, 非第一列, 即有 左边节点 & 上边节点, 取 2 节点中路径和小的
				if dp[i][j-1] < dp[i-1][j] {
					prev = dp[i][j-1]
				} else {
					prev = dp[i-1][j]
				}
			}

			dp[i][j] = prev + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
