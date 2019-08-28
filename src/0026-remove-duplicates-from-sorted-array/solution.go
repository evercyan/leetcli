package solution

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for _, v := range nums {
		if v != nums[i] {
			i++
			nums[i] = v
		}
	}
	nums = nums[:i+1]
	return i + 1
}
