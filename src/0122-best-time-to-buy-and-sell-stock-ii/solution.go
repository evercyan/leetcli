package solution

/**
 * [1, 2, 4]
 * 卖出和买入可以在同一天,
 * 1. 在 1 买入, 在 4 卖出, p = 4 - 1 = 3
 * 2. 在 1 买入, 在 2 卖出, p1 = 2 - 1 = 1, 紧接着在 2 买入, 在 4 卖出, p2 = 4 - 2 = 2, p = p1 + p2 = 3
 */

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}
