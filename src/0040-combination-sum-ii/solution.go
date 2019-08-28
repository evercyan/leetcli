package solution

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	set := []int{}
	sum, pos := 0, 0
	handle(candidates, set, target, sum, pos, &ret)
	return ret
}

func handle(candidates, set []int, target, sum, pos int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, set...))
		return
	}
	// 从起始位置循环, 循环条件满足当前和 + 当前元素 <= 目标值
	for i := pos; i < len(candidates) && (sum+candidates[i]) <= target; i++ {
		if i != pos && candidates[i] == candidates[i-1] {
			continue
		}
		handle(candidates, append(set, candidates[i]), target, sum+candidates[i], i+1, result)
	}
}
