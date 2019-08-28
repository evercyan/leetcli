package solution

/**
 * Attention
 *
 * 此题有个限制条件是数组元素连续相加, 而不是随机组合, 故, 单循环即可
 * 不然会有点麻烦~
 */

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	tmpsum, num := 0, 0
	for i := 0; i < len(A); i++ {
		if tmpsum == sum/3 {
			num++
			tmpsum = 0
		}
		tmpsum += A[i]
		if num >= 2 {
			break
		}
	}
	return num >= 2
}
