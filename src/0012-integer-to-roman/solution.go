package solution

/**
 * map range 无序
 */
func intToRoman(num int) string {
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strings := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	s := ""
	for i := 0; i < len(ints); i++ {
		for {
			if num < ints[i] {
				break
			}
			num -= ints[i]
			s += strings[i]
		}
	}
	return s
}
