package solution

func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= 1 {
		return s
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += 2*numRows - 2 {
			ret += string(s[j])
			if i != 0 && i != numRows-1 && j+2*(numRows-i-1) < len(s) {
				ret += string(s[j+2*(numRows-i-1)])
			}
		}
	}
	return ret
}
