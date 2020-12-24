package solution

// 二分法查找

func isBadVersion(num int) bool {
	// 此处 4 为预设, 和 solution_test.go 中 expects 4 对应
	return num >= 4
}

func firstBadVersion(n int) int {
	low, high := 1, n
	for low < high {
		mid := (low + high) / 2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
