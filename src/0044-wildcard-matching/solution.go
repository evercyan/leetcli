package solution

/**
 * 此题型类似 0010-regular-expression-matching, 同样用动态规则处理
 *
 * dp[i][j] 表示 s[:i] 与 p[:j] 是否匹配, 最终取 dp[len(s)][len(p)]
 */
func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)
	dp := make([][]bool, sl+1)
	for i := 0; i <= sl; i++ {
		dp[i] = make([]bool, pl+1)
	}

	// 表示空串是匹配的
	dp[0][0] = true

	for i := 1; i <= pl; i++ {
		if p[i-1] != '*' {
			break
		}
		// 处理匹配串 p 以若干星号开头的情况, s 为空串也能匹配 p[:i]
		dp[0][i] = true
	}

	for i := 1; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[sl][pl]
}
