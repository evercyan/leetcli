package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func s2i(v string) int {
	r, _ := strconv.Atoi(v)
	return r
}

func i2s(v int) string {
	return fmt.Sprintf("%v", v)
}

func addStrings(num1 string, num2 string) string {
	// 补全前置位 0, 使位数对齐
	l1, l2 := len(num1), len(num2)
	if l1 > l2 {
		num2 = strings.Repeat("0", l1-l2) + num2
	} else {
		num1 = strings.Repeat("0", l2-l1) + num1
	}

	ret := ""
	in := 0
	for i := len(num1) - 1; i >= 0; i-- {
		tmp := s2i(string(num1[i])) + s2i(string(num2[i])) + in
		in = tmp / 10
		ret = i2s(tmp%10) + ret
	}
	if in > 0 {
		ret = "1" + ret
	}

	return ret
}
