package solution

func findSpecialInteger(arr []int) int {
	len := len(arr)
	if len == 1 {
		return arr[0]
	}
	// 纪录当前数
	num := arr[0]
	// 纪录当前数出现次数
	count := 1
	for i := 1; i < len; i++ {
		// 判断之
		if count*4 > len {
			return num
		}
		if num == arr[i] {
			// 相等, 则出现数++
			count++
		} else {
			// 改变当前数
			num = arr[i]
			count = 1
		}
	}
	return num
}
