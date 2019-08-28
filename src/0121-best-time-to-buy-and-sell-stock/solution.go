package solution

/**
 * 简单循环处理
 */
/*
func maxProfitR(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
*/

/**
 * 动态规划处理
 * 将 maxProfitR2 抽象化为:
 * 前 i 天的最大收益 = max(前 i-1 天的最大收益, 第 i 天的价格 - 前 i-1 天中的最小价格)
 */
/*
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxProfitDP(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// 纪录第 N 天的最大利润
	data := make([]int, len(prices))
	data[0] = 0
	// 纪录前 N 天的最小价格
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		// 计算第 N 天的最大收益
		data[i] = max(prices[i]-min, data[i-1])
		if prices[i] < min {
			min = prices[i]
		}
	}
	max := 0
	// 遍历天数, 找出最大收益
	for _, p := range data {
		if p > max {
			max = p
		}
	}
	return max
}
*/
/**
 * 动态规划优化
 */
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// 纪录最小价格
	max, minP := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minP {
			// 当某天价格小于最小价格时, 重赋值
			minP = prices[i]
			continue
		}
		if prices[i]-minP > max {
			// 判断某天价格差
			max = prices[i] - minP
		}
	}
	return max
}
