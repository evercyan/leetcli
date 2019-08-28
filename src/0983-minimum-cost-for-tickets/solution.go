package solution

/**
 * 动态规范
 *
 * dp[i] 表示从开始到第 i 天的最少消费
 * 如果第 i 天不在 days 里, 直接 dp[i-1]
 * 或者:
 * dp[i] = min(dp[i-1] + c0, dp[i-7] + c1, dp[i-30] + c2)
 * min(i-1 最低消费 + 日票, i-7 最低消费 + c2, i-30 最低消费 + c3)
 */

func min(x int, y int) int {
	min := x
	if min > y {
		min = y
	}
	return min
}

func max(x int, y int) int {
	max := x
	if max < y {
		max = y
	}
	return max
}

func mincostTickets(days []int, costs []int) int {
	if len(days) <= 0 {
		return 0
	}
	dp := make([]int, days[len(days)-1]+1)
	dp[0] = 0
	dMap := map[int]int{}
	for _, d := range days {
		dMap[d] = 1
	}
	for i := 1; i <= days[len(days)-1]; i++ {
		if _, ok := dMap[i]; !ok {
			// 如果当天不在旅行日里, 花费延用 i-1
			dp[i] = dp[i-1]
			continue
		}
		// max(0, i-7), max(0, i-30) 防止数组越界, 如果没有前 7 天, 从第一天开始
		dp[i] = min(dp[i-1]+costs[0], dp[max(0, i-7)]+costs[1])
		dp[i] = min(dp[i], dp[max(0, i-30)]+costs[2])
	}
	return dp[days[len(days)-1]]
}
