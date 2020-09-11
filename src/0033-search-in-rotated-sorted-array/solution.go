package solution

/**
 * 你的算法时间复杂度必须是 O(log n) 级别
 *
 * 1. 先判断数组是否是递增数组, 是, 直接二分法查找
 * 2. 不是递增加, 取 mid = len / 2, 递规则左边数组和右边数组, 重复 1
 * 3. 注意, 如果是右边数组返回满足的 key, 需要加上偏移的 mid
 */

func search(nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return -1
	}
	if nlen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if nums[0] < nums[nlen-1] {
		// 有序数组, 二分法查找
		l, h := 0, nlen-1
		ret := -1
		for l <= h {
			mid := (l + h) / 2
			if nums[mid] == target {
				ret = mid
				break
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
		return ret
	}
	// 无序数组, 一分为二, 递规左右两个子数组
	mid := nlen / 2
	lRet := search(nums[:mid], target)
	if lRet >= 0 {
		return lRet
	}
	rRet := search(nums[mid:], target)
	if rRet >= 0 {
		// key 加上偏移值
		return rRet + mid
	}
	return -1
}
