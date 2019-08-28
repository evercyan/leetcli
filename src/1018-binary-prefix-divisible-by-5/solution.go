package solution

func prefixesDivBy5(A []int) []bool {
	ret := []bool{}
	if len(A) <= 0 {
		return ret
	}
	num := 0
	for i := 0; i < len(A); i++ {
		if i == 0 {
			num = A[i]
		} else {
			num = num<<1 + A[i]
		}
		ret = append(ret, num%5 == 0)
		// 此处需补足取余操作, 否则 num 值会溢出
		// 因为只要个位数 0, 5 都满足, 所以只需要保留个位数即可
		num = num % 10
	}
	return ret
}
