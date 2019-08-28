package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret = [][]int{}
	nlen := len(nums)
	if nlen < 3 {
		return ret
	}
	// 先排序, 方便相同值跳过
	sort.Ints(nums)
	for i := 0; i < nlen-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := nlen - 1
		// 得到 j + k 的值
		hit := 0 - nums[i]
		// 循环处理
		for j < k {
			tmpHit := nums[j] + nums[k]
			if tmpHit == hit {
				// 写入结果集
				tmp := []int{nums[i], nums[j], nums[k]}
				ret = append(ret, tmp)
				// 同时 j++, k--
				j++
				k--
				// 跳过相同值
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// 跳过相同值
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if tmpHit > hit {
				// 如果当前值大于目标值, 将 k--
				k--
			} else {
				// 如果当前值小于目标值, 将 k++
				j++
			}
		}
	}
	return ret
}
