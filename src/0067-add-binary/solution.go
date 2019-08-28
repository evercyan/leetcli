package solution

import (
	"strings"
)

/**
 * 题目解读
 *
 * 1. 先补足 a, b 等长
 * 2. 逆序遍历, 在不同进位下, 判断 a["0", "1"], b["0", "1"] 的状态
 * 3. 当遍历结束时, 如果仍有进位, 左位补 "1"
 */

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	ret := ""
	// 进位标识
	flag := 0
	for i := len(a) - 1; i >= 0; i-- {
		char := ""
		if flag == 0 {
			// 无进位情况下
			if a[i:i+1] == "1" && b[i:i+1] == "1" {
				char = "0"
				flag = 1
			} else if a[i:i+1] == "1" || b[i:i+1] == "1" {
				char = "1"
				flag = 0
			} else {
				char = "0"
				flag = 0
			}
		} else {
			// 有情况位情况下
			if a[i:i+1] == "1" && b[i:i+1] == "1" {
				char = "1"
				flag = 1
			} else if a[i:i+1] == "1" || b[i:i+1] == "1" {
				char = "0"
				flag = 1
			} else {
				char = "1"
				flag = 0
			}
		}
		ret = char + ret
		if i == 0 && flag == 1 {
			// 遍历结束后, 如果仍有进位
			ret = "1" + ret
		}
	}
	return ret
}
