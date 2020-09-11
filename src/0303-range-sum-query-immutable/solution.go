package solution

/**
 * 题解
 *
 * 会多次调用 sumRange 方法, 故简单的循环不可取
 * 走动态规划备忘录法
 *
 * dp[i] = dp[i-1] + num[i]
 *
 * Attention: SumRange(i, j) 包含 i, j, 故
 * SumRange(i, j) = dp[j] - dp[i-1](注意此处不是 dp[i], 是 dp[i-1])
 */

type NumArray struct {
	nums []int
	dp   []int
}

func Constructor(nums []int) NumArray {
	if len(nums) <= 0 {
		return NumArray{
			nums: []int{},
			dp:   []int{},
		}
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
	}
	return NumArray{
		nums: nums,
		dp:   dp,
	}
}

func (na *NumArray) SumRange(i int, j int) int {
	if j > len(na.nums) || i < 0 {
		return 0
	}
	if i == 0 {
		return na.dp[j]
	}
	return na.dp[j] - na.dp[i-1]
}
