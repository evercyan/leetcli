package solution

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	maxLen := 1
	start := 0
	for i := 0; i < sLen; i++ {
		if i-maxLen >= 1 && s[i-maxLen-1:i+1] == reverseString(s[i-maxLen-1:i+1]) {
			start = i - maxLen - 1
			maxLen += 2
			continue
		}
		if i-maxLen >= 0 && s[i-maxLen:i+1] == reverseString(s[i-maxLen:i+1]) {
			start = i - maxLen
			maxLen += 1
		}
	}
	return s[start : start+maxLen]
}
