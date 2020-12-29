package solution

func firstUniqChar(s string) int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] == 1 {
			return i
		}
	}
	return -1
}
