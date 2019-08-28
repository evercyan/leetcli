package solution

/**
 * 如何用动态规划来做
 */

func max(x int, y int) int {
	max := x
	if max < y {
		max = y
	}
	return max
}

func maxTurbulenceSize(A []int) int {
	if len(A) <= 0 {
		return 0
	}
	// 结果值, 符合条件游标
	ret, cursor := 1, 0
	for i := 1; i < len(A); i++ {
		// 如果有前后相等的情况
		if A[i] == A[i-1] {
			cursor = i
			continue
		}
		// 符合条件, 或者循环到数组结束
		if i == len(A)-1 || (A[i-1] <= A[i] && A[i] <= A[i+1]) || (A[i-1] >= A[i] && A[i] >= A[i+1]) {
			// 比较之
			ret = max(ret, i-cursor+1)
			// 游标重置
			cursor = i
		}
	}
	return ret
}
