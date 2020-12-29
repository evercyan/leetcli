package solution

import (
	"strings"
)

func decodeString(s string) string {
	ret := ""

	// 次数栈和字符栈单独存储
	numStack, strStack := []int{}, []string{}
	num := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			// 处理多位连续数字, e.g. 12[a]
			num = num*10 + int(char-'0')
		} else if char == '[' {
			// 先将次数压入栈
			numStack = append(numStack, num)
			num = 0
			// 将前置字符串压入栈, 此后的 ret 直到 ] 前都是待重复字符
			strStack = append(strStack, ret)
			ret = ""
		} else if char == ']' {
			// 取重复次数
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			// 取前置字符串
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			ret = str + strings.Repeat(ret, count)
		} else {
			ret += string(char)
		}
	}

	return ret
}
