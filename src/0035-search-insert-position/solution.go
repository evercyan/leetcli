package solution

func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	// nums[i] == target, 返回索引 i
	// nums[i] > target, target 替换掉 nums[i] 所在的位置, 返回 i
	// nums[i] < target, i++, 如果遍历结束仍未有符合上面两种 case, 则插入位置为数组末尾, 返回索引 len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
