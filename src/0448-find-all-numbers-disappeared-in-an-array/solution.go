package solution

/**
 * 普通解法, 走 hash map 保存判断
 * 1080ms > 69.77%
 * 15.3 MB > 39.29%
 */
func findDisappearedNumbers(nums []int) []int {
	ret := []int{}
	if len(nums) <= 0 {
		return ret
	}
	numMap := make(map[int]int)
	for i := 1; i <= len(nums); i++ {
		numMap[i] = 0
	}
	for _, num := range nums {
		numMap[num]++
	}
	for k, v := range numMap {
		if v == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}

/**
 * 交换位置, 让数字回到它到本来的位置
 * 860ms > 91.86%
 * 22.1 MB > 7.14%
 */
/*
func findDisappearedNumbers(nums []int) []int {
	ret := []int{}
	if len(nums) <= 0 {
		return ret
	}
	for i := 0; i < len(nums); {
		if nums[i] == i+1 || nums[i] == nums[nums[i]-1] {
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
*/
