package solution

/**
 * 题目解读
 *
 * 1. 遍历字符串, 如果为 "(", 写入临时数组
 * 如果为 ")", 从临时数组 pop 出最后一个元素的 key, 赋值为 1
 * 2. 找到最大连续 1 的数量
 */

func longestValidParentheses(s string) int {
	list, tmp := make([]int, len(s)), []int{}
	for k, v := range s {
		if string(v) == "(" {
			tmp = append(tmp, k)
			continue
		}
		if len(tmp) > 0 {
			// 相当于 tmp.pop()
			list[tmp[len(tmp)-1]], list[k] = 1, 1
			tmp = append([]int{}, tmp[:len(tmp)-1]...)
		}
	}
	// 循环找连续为 1 的最大数量
	ret, tmpret := 0, 0
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			tmpret = 0
			continue
		}
		tmpret++
		if tmpret > ret {
			ret = tmpret
		}
	}
	return ret
}
