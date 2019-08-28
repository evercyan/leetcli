package solution

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	// 先排序
	sort.Ints(candidates)
	// 保存单次结果
	set := []int{}
	// 循环起始位置
	pos := 0
	handle(candidates, target, pos, &ret, set)
	return ret
}

func handle(candidates []int, target int, pos int, ret *[][]int, set []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// 复制, 不影响原 set 值
		c := make([]int, len(set))
		copy(c, set)
		*ret = append(*ret, c)
		return
	}
	for i := pos; i < len(candidates); i++ {
		set = append(set, candidates[i])
		handle(candidates, target-candidates[i], i, ret, set)
		set = set[:len(set)-1]
	}
}
