package solution

// import (
// 	"strings"
// )

/**
 * 1. 标准解法 O(m*n)
 */

func longestCommonPrefix(strs []string) string {
	sLen := len(strs)
	if sLen == 0 {
		return ""
	} else if sLen == 1 {
		return strs[0]
	}
	// 找到所有字符串中最小长度
	minStrLen := len(strs[0])
	for i := 1; i < sLen; i++ {
		if len(strs[i]) < minStrLen {
			minStrLen = len(strs[i])
		}
	}
	// 初始索引
	index := minStrLen - 1
	// 循环每一个字符串
	for i := 1; i < sLen; i++ {
		// 循环长度, 如果不匹配, 索引 = 当前位 - 1, 继续下一个字符串比较
		for j := 0; j <= index; j++ {
			if strs[i][j] != strs[0][j] {
				index = j - 1
				break
			}
		}
	}
	if index < 0 {
		return ""
	}
	return strs[0][0 : index+1]
}

/**
 * 2. 二分法 O(log2n * m)
 */
/*
func isCommonPrefix(strs []string, prefix string) bool {
	for _, str := range strs {
		if !strings.HasPrefix(str, prefix) {
			return false
		}
	}
	return true
}
func longestCommonPrefix(strs []string) string {
	sLen := len(strs)
	if sLen == 0 {
		return ""
	} else if sLen == 1 {
		return strs[0]
	}
	minStrLen := len(strs[0])
	for i := 1; i < sLen; i++ {
		if len(strs[i]) < minStrLen {
			minStrLen = len(strs[i])
		}
	}
	low, high := 0, minStrLen
	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, strs[0][0:mid]) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if high <= 0 {
		return ""
	}
	return strs[0][0:high]
}
*/
