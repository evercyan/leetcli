package solution

import (
	"sort"
)

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func threeSumClosest(nums []int, target int) int {
	// 先排序数组, 从小到大
	sort.Ints(nums)
	// 先定义个初始最小值
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		// 双指针查找
		l, h := i+1, len(nums)-1
		for l < h {
			num := nums[i] + nums[l] + nums[h]
			if num > target {
				h--
			} else if num < target {
				l++
			} else {
				return target
			}
			// 比较差值绝对值, 找逼近数
			if abs(num, target) < abs(closest, target) {
				closest = num
			}
		}
	}
	return closest
}
