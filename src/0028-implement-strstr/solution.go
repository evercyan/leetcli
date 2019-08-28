package solution

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	// key, 记录 needle 游标;
	// index, 记录匹配时 haystack 的游标;
	key, index := 0, 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[key] {
			// 在匹配不相等时, 如果目标 key 已经偏移, 需要将 i 回溯
			if key != 0 {
				i = i - key
			}
			key = 0
			continue
		}
		// 匹配情况下, 如果 key 未偏移, 则纪录 index
		if key == 0 {
			index = i
		}
		key++
		if key == len(needle) {
			return index
		}
	}
	return -1
}
