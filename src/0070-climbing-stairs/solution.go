package solution

/**
 * 题目解读
 *
 * 1. 斐波那契数列
 */

func climbStairs(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
