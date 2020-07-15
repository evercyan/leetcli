package solution

/**
 * 思路:
 * 蚂蚁碰撞调头相当于碰撞后若无其事的继续走, 故此题直接求左右到终点的最大长度即可
 */

func getLastMoment(n int, left []int, right []int) int {

	max := 0

	for i := 0; i < len(left); i++ {
		step := left[i] - 0
		if step > max {
			max = step
		}
	}

	for i := 0; i < len(right); i++ {
		step := n - right[i]
		if step > max {
			max = step
		}
	}

	return max
}
