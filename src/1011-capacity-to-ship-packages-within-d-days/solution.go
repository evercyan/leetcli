package solution

/**
 * 题解
 *
 * 1. 货物连续
 * 2. 每次打包 D--, 保证 D >= 0 即可
 *
 * 使用二分搜索
 * 结果 r, 满足 max(weights) <= r <= sum(weights)
 */

func shipWithinDays(weights []int, D int) int {
	// 先算出 r 的上下限
	low, high := 0, 0
	for _, w := range weights {
		high += w
		if w > low {
			low = w
		}
	}
	for low < high {
		// r 取中间值, 不断逼近
		mid := (low + high) / 2
		// 运输总值, 天数
		w, d := 0, 0
		for i := 0; i < len(weights); i++ {
			// 如果 + 下一个超过预设值, d++, 且重置 w,
			if w+weights[i] > mid {
				d++
				w = 0
			}
			// 运输 +
			w += weights[i]

			// 数组执行结束, 补足 d++
			if i == len(weights)-1 {
				d++
			}
		}
		if d <= D {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
