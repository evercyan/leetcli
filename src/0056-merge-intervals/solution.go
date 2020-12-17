package solution

import (
	"sort"
)

// 遍历, 取前后节点比较即可

func merge(intervals [][]int) [][]int {
	var ret [][]int

	// 按第 1 位升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] < intervals[i][0] {
			// 前后并不相连, prev 推入结果集
			ret = append(ret, prev)
			// 变换 prev, 继续处理
			prev = intervals[i]
		} else {
			// 前后相连, 比对 prev[1] 和 intervals[i][1] 大小, 决定结束位置
			if prev[1] < intervals[i][1] {
				prev[1] = intervals[i][1]
			}
		}
	}

	// 最后补入
	ret = append(ret, prev)

	return ret
}
