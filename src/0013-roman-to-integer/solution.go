package solution

func romanToInt(s string) int {
	roman := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	sLen := len(s)
	sum := 0
	for i := 0; i < sLen-1; i++ {
		p, _ := roman[string(s[i])]
		q, _ := roman[string(s[i+1])]
		if p < q {
			sum -= roman[string(s[i])]
		} else {
			sum += roman[string(s[i])]
		}
	}
	sum += roman[string(s[sLen-1:sLen])]
	return sum
}
