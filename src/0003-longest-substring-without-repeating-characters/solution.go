package solution

func lengthOfLongestSubstring(s string) int {
	start, max := 0, 0
	var usedChar = make(map[string]int)
	for i := range s {
		char := string(s[i])
		if _, ok := usedChar[char]; ok && start <= usedChar[char] {
			start = usedChar[char] + 1
		} else {
			if max < i-start+1 {
				max = i - start + 1
			}
		}
		usedChar[char] = i
	}
	return max
}
