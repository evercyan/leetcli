package solution

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	p := x
	q := 0
	for p != 0 {
		q = q*10 + p%10
		p /= 10
	}
	return q == x
}
