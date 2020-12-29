package solution

func getRow(rowIndex int) []int {
	// 上一次结果
	ret := []int{1}

	for i := 1; i < rowIndex+1; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
			} else {
				tmp[j] = ret[j-1] + ret[j]
			}
		}
		ret = tmp
	}

	return ret
}
