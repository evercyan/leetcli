package solution

/**
 * 你的算法时间复杂度必须是 O(log n) 级别
 *
 * 二分法查找, 分别查找起点, 终点
 */

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	nlen := len(nums)
	if nlen == 0 {
		return ret
	}
	// 查找起点
	bl, bh := 0, nlen-1
	for bl <= bh {
		bmid := (bl + bh) / 2
		if nums[bmid] < target {
			bl = bmid + 1
		} else if nums[bmid] > target {
			bh = bmid - 1
		} else {
			// nums[mid] == target
			// 如果当前元素下标是 0, 或者前一个元素不等于当前元素, 当前元素为起点
			if bmid == 0 || (bmid-1 >= 0 && nums[bmid-1] != target) {
				ret[0] = bmid
				break
			}
			bh = bmid - 1
		}
	}
	// 查找终点
	el, eh := 0, nlen-1
	for el <= eh {
		emid := (el + eh) / 2
		if nums[emid] < target {
			el = emid + 1
		} else if nums[emid] > target {
			eh = emid - 1
		} else {
			// nums[mid] == target
			// 如果当前元素下标是 nlen-1, 或者后一个元素不等于当前元素, 当前元素为终点
			if emid == nlen-1 || (emid+1 <= nlen && nums[emid+1] != target) {
				ret[1] = emid
				break
			}
			el = emid + 1
		}
	}
	return ret
}
