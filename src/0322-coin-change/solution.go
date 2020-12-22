package solution

// 动态规划
// 令 dp[i] 表示组成金额 i 的最小组合数
// 假设有硬币 c, 则必然有 dp[i+c] = dp[i] + 1
// 即 dp[i]  = dp[i-c]+1

func coinChange(coins []int, amount int) int {
	// 初始化 dp, 有 dp[0]: 0
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			// i < c, 即硬币 c 面额大于总额 i, 无法组合
			if i < c {
				continue
			}

			// dp[i-c] == -1, 即不存在 i-c 总额的最佳组合
			if dp[i-c] == -1 {
				continue
			}

			// 存在 i-c 的最佳组合, 则 dp[i] = dp[i-c] + 1
			// dp[i] 存在且小于基于 i-c 的最佳组合数
			if dp[i] != -1 && dp[i] <= dp[i-c]+1 {
				continue
			}

			dp[i] = dp[i-c] + 1
		}
	}

	return dp[amount]
}
