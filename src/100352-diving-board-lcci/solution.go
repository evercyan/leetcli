package solution

func divingBoard(shorter int, longer int, k int) []int {
	result := []int{}
	if k == 0 {
		return result
	} else if shorter == longer {
		// 须考虑相同长度的情况
		return append(result, shorter*k)
	}

	for i := 0; i <= k; i++ {
		length := longer*i + shorter*(k-i)
		result = append(result, length)
	}

	return result
}
