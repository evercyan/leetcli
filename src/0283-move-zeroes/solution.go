package solution

// 双指针

func moveZeroes(nums []int) {
	// 将所有非 0 的前置
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[index] = nums[i]
		index++
	}

	// 后位全置为 0
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
