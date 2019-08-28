package solution

func reverse(num int) int {
	ret := 0
	for num != 0 {
		ret = ret*10 + num%10
		num /= 10
	}
	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}
