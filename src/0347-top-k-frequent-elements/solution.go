package solution

// todo 堆

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var n []int
	for num, _ := range m {
		n = append(n, num)
	}

	sort.Slice(n, func(i, j int) bool {
		return m[n[i]] > m[n[j]]
	})

	return n[:k]
}
