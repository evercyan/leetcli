package solution

/**
 * 回溯算法
 *
 * l: 左括号数量; r, 右括号数量;
 *
 * 限制条件
 *  l > 0, 才能放入左括号
 *  r > l, 才能放入右括号
 */

func generateParenthesis(n int) []string {
	return dp("", n, n)
}

func dp(s string, l int, r int) []string {
	if l == 0 && r == 0 {
		return []string{s}
	}
	ret := []string{}
	if l > 0 {
		ret = append(ret, dp(s+"(", l-1, r)...)
	}
	if r > l {
		ret = append(ret, dp(s+")", l, r-1)...)
	}
	return ret
}
