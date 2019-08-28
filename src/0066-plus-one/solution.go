package solution

/**
 * 题目解读
 * 4321 -> []int{4, 3, 2, 1} -> +1 -> 4322 -> []int{4, 3, 2, 2}
 * 1019 -> []int{1, 0, 1, 9} -> +1 -> 1020 -> []int{1, 0, 2, 0}
 */

func plusOne(digits []int) []int {
	// 进位标志
	flag := 0
	// 逆循环
	for i := len(digits) - 1; i >= 0; i-- {
		// 加进位
		ni := digits[i] + flag
		// 如果是最后一个元素, +1
		if i == len(digits)-1 {
			ni += 1
		}
		// 大于 10, 设置进制位
		if ni >= 10 {
			ni = ni - 10
			flag = 1
		} else {
			flag = 0
		}
		digits[i] = ni
		// 如果遍历完第一个元素, 仍然有进位, 数组头部插入
		if i == 0 && flag == 1 {
			// 头部追加
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
