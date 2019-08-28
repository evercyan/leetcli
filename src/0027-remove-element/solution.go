package solution

func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}
	key := 0
	for i := 0; i < numsLen; i++ {
		if nums[i] != val {
			nums[key] = nums[i]
			key++
		}
	}
	return key
}
