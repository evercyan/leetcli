package solution

import (
	"strconv"
	"strings"
)

/**
 * 模拟竖式乘法
 * string -> int 可以利用 rune == int32 来做
 */

/*
    1 2 i
    3 4 j
--------
0 0 0 8
0 0 4
0 0 6
0 3
--------
0 4 0 8
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// 结果最长位数, 99*99 = 9801, 两位数相乘, 积长度最大为两乘数长度相加
	pos := make([]int, len(num1)+len(num2))
	ret := ""
	// 倒叙遍历, 从个位开始
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// rune -> int 实现乘法
			num := int((num1[i] - '0') * (num2[j] - '0'))
			// p1, i, j 相乘后最高的位
			// p2, 表示 p1 的下一位
			p1, p2 := i+j, i+j+1
			// 将当前结果 + 原先 p2 位结果后, 再分别取余处理
			sum := num + pos[p2]
			// p1 已经是最高位, 不用考虑 sum > 100 的情况
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}
	for i := 0; i < len(pos); i++ {
		ret += strconv.Itoa(pos[i])
	}
	return strings.TrimLeft(ret, "0")
}
