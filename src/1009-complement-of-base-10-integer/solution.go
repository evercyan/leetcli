package solution

import (
	"fmt"
)

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	// 计算 N 的二进制
	s := ""
	for i := N; i > 0; i /= 2 {
		s = fmt.Sprintf("%v%v", i%2, s)
	}
	// 反码 = (原码 & 反码的和) - 原码
	// 原码二过制长度为 3, 则和为 2(^3) - 1
	sum := 1
	for i := 0; i < len(s); i++ {
		sum *= 2
	}
	sum -= 1
	return sum - N
}
