package solution

/**
 * 普通解法, O(n2)
 * 3092 ms > 52.08%
 * 14.3 MB > 77.27%
 */

func dailyTemperatures(T []int) []int {
	n := len(T)
	ret := make([]int, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				ret[i] = j - i
				break
			}
		}
	}
	return ret
}

/**
 * 模拟栈
 * 循环从栈中 pop 出数据与当前节点比较, 因为 i++, 故只要 >, 即最近
 */
/*
func dailyTemperatures(T []int) []int {
	n := len(T)
	ret, stack := make([]int, n), make([]int, n)
	top := -1
	for i := 0; i < n; i++ {
		for top >= 0 && T[stack[top]] < T[i] {
			ret[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}
	return ret
}
*/
