package solution

import (
	"strconv"
)

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	s := ""
	for i := 0; N > 0; i++ {
		s = strconv.FormatInt(int64(N%2), 10) + s
		// Attention: 奇位需要补足 1
		if i%2 == 1 {
			N += 1
		}
		N /= 2
	}
	return s
}
