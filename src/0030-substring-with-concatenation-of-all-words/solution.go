package solution

func key(list []string, target string) int {
	for k, v := range list {
		if v == target {
			return k
		}
	}
	return -1
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	// 结果集
	ret := []int{}
	// 单词长度
	wlen := len(words[0])
	// 单词数量
	wcount := len(words)
	// i 从开开始, 因为要连续, 故最大偏移量为 字符串长度 - 单词数量*单词长度 + 1
	for i := 0; i < len(s)-wlen*wcount+1; i++ {
		count := 0
		// 数组 copy, map 值引用
		list := make([]string, wcount)
		copy(list, words)
		// j 不是下标, 而是 slice 切片, 故最大到 len(s)
		for j := i; j+wlen <= len(s); j += wlen {
			if count >= wcount {
				break
			}
			key := key(list, s[j:j+wlen])
			if key != -1 {
				list = append(list[:key], list[key+1:]...)
				count++
			} else {
				break
			}
		}
		if count == wcount {
			ret = append(ret, i)
		}
	}
	return ret
}
