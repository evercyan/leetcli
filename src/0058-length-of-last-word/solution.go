package solution

import (
	"strings"
)

/**
 * trim 掉字符串末尾所有空格
 */
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	if s == "" {
		return 0
	}
	sList := strings.Split(s, " ")
	return len(sList[len(sList)-1])
}
