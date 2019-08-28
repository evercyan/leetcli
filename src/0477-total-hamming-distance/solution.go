package solution

/**
 * O(n2)
 * 遍历数组, & 1 后得到最后位值 0 || 1
 * 记录 0 的次数 p  & 1 的次数 q, 最后位汉明距离必为 p*q
 * 当前元素向右移位 1
 * 直至所有的元素均右移为 0 后, 停止循环
 */
func totalHammingDistance(nums []int) int {
	r := 0
	bit := [2]int{0, 0}
	zeroCount := 0
	numsLen := len(nums)
	for zeroCount < numsLen {
		zeroCount = 0
		bit[0], bit[1] = 0, 0
		for i := 0; i < numsLen; i++ {
			bit[nums[i]&1]++
			if nums[i] == 0 {
				zeroCount++
			} else {
				nums[i] = nums[i] >> 1
			}
		}
		r += bit[0] * bit[1]
	}
	return r
}
