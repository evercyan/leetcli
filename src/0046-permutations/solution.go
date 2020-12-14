package solution

// dfs 遍历 + 回溯算法

func permute(nums []int) [][]int {
	var res [][]int
	// 单独纪录已使用过的数
	var used = make(map[int]bool, len(nums))

	var dfs func(paths []int)
	dfs = func(paths []int) {
		if len(paths) == len(nums) {
			res = append(res, paths)
			return
		}
		for _, num := range nums {
			if used[num] {
				continue
			}
			// 做出选择
			used[num] = true

			// 回溯处理
			// 此处 append(paths, num) 产生新变量, 不影响 line:13 中的 append
			dfs(append(paths, num))

			// 撤销选择, 是为了切入下一个分支处理
			used[num] = false
		}
	}
	dfs([]int{})

	return res
}
