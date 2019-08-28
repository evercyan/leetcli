package solution

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var r = [][]int{}
	nlen := len(nums)
	if nlen < 4 {
		return r
	}
	sort.Ints(nums)
	for i := 0; i < nlen-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmpI := target - nums[i]
		for j := i + 1; j < nlen-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			tmpJ := tmpI - nums[j]
			for k := j + 1; k < nlen-1; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				tmpK := tmpJ - nums[k]
				for l := k + 1; l < nlen; l++ {
					if nums[l] == tmpK {
						tmp := []int{nums[i], nums[j], nums[k], nums[l]}
						r = append(r, tmp)
						break
					}
				}
			}
		}

	}
	return r
}
