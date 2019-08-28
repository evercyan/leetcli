package solution

/**
 * 动态规划
 *
 * dp[i][j] 表示 s[i:] 与 p[j:] 是否匹配, 最终取 dp[0][0]
 *
 * 1. 判断 s[i] == p[j] || p[j] == ".", 也就是 s[i] 和 p[j] 是否匹配, 结果记 match
 * 2. 存在 p[j+1] 且 p[j+1] == "*", 即 p[j] 可出现 0-n 次
 * 2.1 match == false, 跳过 p[j] 和 p[j+1], 此时, dp[i][j] = dp[i][j+2](跳过两个字符再比较)
 * 2.2 match == true, 则 p[j] 和 p[j+1] 符合 a* 或 .*, i+1 再比, 此时 dp[i][j] = dp[i+1][j]
 * 3. 不存在 p[j+1] == "*"
 * 3.1 match == false, dp[i][j] = false
 * 3.2 match == true, dp[i][j] = dp[i+1][j+1](如果 dp[i+1][j+1] 匹配, 则 dp[i][j] 匹配)
 */

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			match := false
			if i < len(s) && (s[i] == p[j] || p[j] == '.') {
				match = true
			}
			// 当第二个字符模式中有 * 时
			if (j+1) < len(p) && p[j+1] == '*' {
				// dp[i][j+2] 是考虑 * 代表 0 时的情况
				// match && dp[i+1][j] 是考虑在 ij 匹配情况下, 存在 .* 或 i*, 偏移至 i+1
				dp[i][j] = dp[i][j+2] || (match && dp[i+1][j])
			} else {
				dp[i][j] = match && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
