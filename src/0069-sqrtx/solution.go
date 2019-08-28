package solution

import (
	"math"
)

/**
 * 题目解读
 *
 * 牛顿法求平方根
 * x = (x + n/x) / 2
 */

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	ret := x
	for ret > x/ret {
		ret = (ret + x/ret) / 2
	}
	return int(math.Floor(float64(ret)))
}
