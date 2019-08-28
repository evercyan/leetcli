package solution

/**
 * 题解, 所有文字我都看得懂, 连起来我就看不懂了...
 * 参考热评的题目解读
 *
 * 给定一个数组
 * 1, 从右向左遍历, 找到第一个 i < i+1 的 i 节点
 * 2, 如果不存在 i, 反转整个数组
 * 3, 如果存在 i, 数组从右向左遍历, 找到第一个 j 使得 nums[j] > nums[i]
 * 4, 交换数组的 i, j
 * 5, 再将 i 后面的元素反转
 */

func reverse(nums []int, x int, y int) {
	for ; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}

func nextPermutation(nums []int) {
	nlen := len(nums)
	// 从右向左遍历, 查找第一个 key < key+1 的 key 节点
	key := nlen - 2
	for ; key >= 0; key-- {
		if nums[key] < nums[key+1] {
			break
		}
	}
	// 如果不存在这样的 key, 反转数组
	if key < 0 {
		reverse(nums, 0, nlen-1)
		return
	}
	// 存在, 从右向左遍历, 找到第一个大于 nums[key] 的数, 交换之
	for i := nlen - 1; i > key; i-- {
		if nums[i] > nums[key] {
			nums[i], nums[key] = nums[key], nums[i]
			break
		}
	}
	// 再将 key 后面的元素反转
	reverse(nums, key+1, nlen-1)
	return
}
