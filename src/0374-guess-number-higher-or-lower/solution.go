package solution

// 二分法查找

func guess(num int) int {
	// 此处 32 为预设, 和 solution_test.go 中 expects 32 对应
	if num == 32 {
		return 0
	} else if num > 32 {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	low, high := 1, n
	for low < high {
		mid := (low + high) / 2
		ret := guess(mid)
		if ret == 0 {
			return mid
		} else if ret == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

/*

   low = 0
   high = n
   while low <= high:
       mid = (low + high) // 2
       guessret = guess(mid)
       if guessret == 0:
           return mid
       elif guessret == -1:
           high = mid - 1
       else:
           low = mid + 1
   return high
*/
