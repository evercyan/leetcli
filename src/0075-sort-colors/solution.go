package solution

// 1. 普通排序
// 2. 双指针交换

func sortColors(nums []int) {
	start, end := 0, len(nums)-1
	for i := 0; i <= end; i++ {
		if nums[i] == 0 {
			// 遇 0 交换, 指针后移
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
		if nums[i] == 2 {
			// 遇 2 移至末尾, 指针前移, 且遍历指针 i--, 继续对刚交换的元素进行处理
			nums[end], nums[i] = nums[i], nums[end]
			i--
			end--
		}
	}
}

// end: 2

// 2 0 1
// 1 0 2
