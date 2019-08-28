package solution

/**
 * 解题思路
 *
 * 1. 先遍历数组, 找出所有 1 的索引, 判断 1 的数量如果不为 3 的倍数, false
 * 2. 如果 1 的数量为 0, 返回 0:len(A)-1
 * 3. 根据最后一个 1 的索引值到数组的长度, 得出后缀补足 0 的数量
 * 4. 将 1 索引数组长度 3 等分, 配合后缀补足 0 的数量, 确认两个偏移 p, q
 * 5. 遍历数组, 根据 p, q 计算 3 个等分的值, 比较之
 */

func threeEqualParts(A []int) []int {
	ret := []int{-1, -1}
	// 存储 1 的位置
	iMap := []int{}
	for k, v := range A {
		if v == 1 {
			iMap = append(iMap, k)
		}
	}
	// 如果全是 0, 也符合条件
	if len(iMap) == 0 {
		ret = []int{0, len(A) - 1}
		return ret
	}
	// 如果 1 的个数不为 3 倍数, 不符合条件
	if len(iMap)%3 != 0 {
		return ret
	}
	// 查找最后一个 1 后面 0 的数量; (数组索引最大值) - (最后一个 1 的索引值)
	suffixZeroNum := (len(A) - 1) - iMap[len(iMap)-1]
	// 每个等分中 1 的数量
	oneNum := len(iMap) / 3
	// 找到三等分中的偏移 p, q, 注意 q 的结尾 + 1, 是因为条件中是 0:i, i+1:j-1, j:len(A)-1
	p, q := iMap[oneNum-1]+suffixZeroNum, iMap[2*oneNum-1]+suffixZeroNum+1
	// 直接转十进制比较, 二进制前缀 0 << 1, 仍为 0,
	n1, n2, n3 := 0, 0, 0
	for i := 0; i < len(A); i++ {
		if i <= p {
			n1 = n1<<1 + A[i]
		} else if i > p && i < q {
			n2 = n2<<1 + A[i]
		} else {
			n3 = n3<<1 + A[i]
		}
	}
	if n1 == n2 && n2 == n3 {
		ret = []int{p, q}
	}
	return ret
}
