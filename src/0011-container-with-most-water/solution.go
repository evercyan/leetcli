package solution

/**
 * 面积: min(i, j) * (j-i)
 */

func min(x, y int) int {
	min := x
	if y < min {
		min = y
	}
	return min
}

func max(x, y int) int {
	max := x
	if y > max {
		max = y
	}
	return max
}

/**
 * 常规解法 0(n2)
 */
/*
func maxArea(height []int) int {
	max, area := 0, 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area = (j - i) * min(height[i], height[j])
			if area > max {
				max = area
			}
		}
	}
	return max
}
*/

/**
 * 官方题解
 */
func maxArea(height []int) int {
	area, l, h := 0, 0, len(height)-1
	for l < h {
		area = max(area, min(height[l], height[h])*(h-l))
		if height[l] < height[h] {
			l++
		} else {
			h--
		}
	}
	return area
}
