package solution

/**
 * 最大周长
 * 1. 先数组从小到大排序
 * 2. 倒叙遍历连续三个数是否满足 i + i1 > i2
 */

import (
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1 - 2; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
