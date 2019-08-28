package solution

/**
 * 近似题: 11
 */

func trap(height []int) int {
	// 结果
	area := 0
	if len(height) <= 2 {
		return area
	}
	// 先找最高点
	maxIndex := 0
	for i, max := 1, height[0]; i < len(height); i++ {
		if height[i] <= max {
			continue
		}
		max, maxIndex = height[i], i
	}
	// 从左边至最高点
	// 以第一个节点为起点遍历, 后比前小, 则可以接到水, 反之重置起点
	for i, tmp := 1, height[0]; i < maxIndex; i++ {
		if height[i] > tmp {
			tmp = height[i]
		} else {
			area += tmp - height[i]
		}
	}
	// 从右边至最高点
	// 以最后一个节点为起点遍历, 前比后小, 则可以接到水, 反之重置起点
	for i, tmp := len(height)-2, height[len(height)-1]; i > maxIndex; i-- {
		if height[i] > tmp {
			tmp = height[i]
		} else {
			area += tmp - height[i]
		}
	}
	return area
}
