package solution

import (
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	len := len(arr)
	if len < 3 {
		return true
	}

	sort.Ints(arr)
	num := arr[1] - arr[0]
	for i := 2; i < len; i++ {
		if arr[i]-arr[i-1] != num {
			return false
		}
	}
	return true
}
