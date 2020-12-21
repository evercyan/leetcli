package solution

import (
	"sort"
)

// 关联: 全排列(https://leetcode-cn.com/problems/permutations/)
// 从第 1 个元素开始遍历, 对每个元素有两种选择, 用或者不用, 再将其代入下一次遍历

func subsets(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var dfs func(paths []int, key int)
	dfs = func(paths []int, key int) {
		if key == len(nums) {
			// 处理到最后元素, 写入结果
			res = append(res, append([]int{}, paths...))
			return

		}
		// 1. 不用它, 继续处理下一个元素
		dfs(paths, key+1)
		// 2. 用它, 再处理下一个元素
		dfs(append(paths, nums[key]), key+1)
	}
	dfs([]int{}, 0)

	return res
}
