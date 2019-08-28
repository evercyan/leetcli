package solution

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	// 最大保存下标 len(nums), 故长度为 +1
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > len(nums) {
			continue
		}
		dp[nums[i]] = 1
	}
	for i := 1; i < len(dp); i++ {
		if dp[i] == 0 {
			return i
		}
	}
	// 如果全符合, 返回数组长度 +1
	return len(nums) + 1
}
