package solution

/**
 * 参考热评, 使用位运算来处理
 *
 * << 1 相当于乘以 2, >> 1 相当于除以 2
 *
 * 可以 dividend >> 2^n, n 最初为 31, 不断减小 n 去试探
 * 当某个 n 满足 dividend >> n >=divisor 时
 * 表示我们找到了一个足够大的数, dividend - 2^n*divisor >= divisor
 * 再拿剩余的减去 divisor 即可
 *
 * 我们可以以 100/3 为例, 初始商 ret = 0
 * n = 6, 100 / 2^6 = 100 / 64 < 3
 * n = 5,100 / 2^5 = 100 / 32 = 3
 * 故, 取 n = 5 时, 被除数为 100 - 3*2^5 = 4, 此时商为 ret += 2^5 = 32
 * 剩余 4, 再进行循环
 * n = 0, 4 / 2^0 > 3
 * 取 n = 0, 被除数为 4 - 3*2^0 = 1, 此时商为 ret += 2^0 = 32 + 1 = 33
 */

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	max := int(^uint32(0) >> 1)
	min := ^max
	if dividend == min && divisor == -1 {
		return max
	}
	// 取绝对值
	t, d := abs(dividend), abs(divisor)
	ret := 0
	// 从最大 31 位开始循环
	for i := 31; i >= 0; i-- {
		if t>>uint(i) >= d {
			// ret = ret + 2^i
			ret += 1 << uint(i)
			// t = t - d * 2^i
			t -= d << uint(i)
		}
	}
	// 异或计算结果符号
	if dividend^divisor < 0 {
		ret = -ret
	}
	return ret
}
