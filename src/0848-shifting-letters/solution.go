package solution

/*
解题思路:
S: "abc"
shifts: [1, 2, 3]

对于 a, 其实移位了 1 + 2 + 3,
对于 b, 其实移位了 2 + 3,
对于 c, 其实移位了 3

所以, 直接从字符串最后一位开始处理移位, 满足:
1. 移位总数 += shifts[i]
2. 移位后字符 = (当前字符相对a的偏移 + 移位总数) % 26 + 字符a的ASCII

a 的 ASCII 码: int('a'), rune
*/

func shiftingLetters(S string, shifts []int) string {
	length := len(shifts)
	if length <= 0 {
		return S
	}
	s, sum := []byte(S), 0
	for i := length - 1; i >= 0; i-- {
		sum += shifts[i]
		s[i] = byte((int(s[i])-int('a')+sum)%26 + int('a'))
	}
	return string(s)
}
