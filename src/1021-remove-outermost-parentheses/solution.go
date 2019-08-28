package solution

func removeOuterParentheses(S string) string {
	ret := ""
	// 左括号数量
	l := 0
	// 注意判断顺序, 先右再加再左
	for i := 0; i < len(S); i++ {
		// 如果是右括号, 说明闭合, l--
		if S[i] == ')' {
			l--
		}
		// 如果 l > 0, 说明尚未闭合, 该字符有效
		if l > 0 {
			ret += string(S[i])
		}
		// 左括号++
		if S[i] == '(' {
			l++
		}
	}
	return ret
}

/*
func removeOuterParentheses(S string) string {
	ret := ""
	// 左, 右括号数量
	l, r := 0, 0
	// 标识括号数量相等
	index := -1
	for i := 0; i < len(S); i++ {
		if l == r {
			// 相等情况, 重置标识
			index = -1
		} else {
			// 如果不相等, 表示有最外括号
			index = i
		}
		if string(S[i]) == "(" {
			l++
		}
		if string(S[i]) == ")" {
			r++
		}
		if r < l && index != -1 {
			ret += string(S[i])
		}
	}
	return ret
}
*/
