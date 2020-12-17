package solution

func findIntegers(num int) int {
	fib := [32]int{1, 2}
	for i := 2; i < 32; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	ans, k, perBit := 0, 30, 0
	for k >= 0 {
		if (num & (1 << uint(k))) > 0 {
			ans += fib[k]
			if perBit > 0 {
				return ans
			}
			perBit = 1
		} else {
			perBit = 0
		}
		k--
	}
	return ans + 1
}
