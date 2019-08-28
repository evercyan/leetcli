package solution

func nextGreaterElements(nums []int) []int {
	ret := []int{}
	if len(nums) == 0 {
		return ret
	}
	if len(nums) == 1 {
		ret = append(ret, -1)
		return ret
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				ret = append(ret, nums[j])
				break
			}
			// 循环至最右, 仍无比当前元素大的, 再从列表的起始到当前元素位置, 循环比对一次
			if j == len(nums)-1 {
				// 如果当前是第一个元素, 不用循环了
				if i == 0 {
					ret = append(ret, -1)
				} else {
					for k := 0; k < i; k++ {
						if nums[k] > nums[i] {
							ret = append(ret, nums[k])
							break
						}
						if k == i-1 {
							ret = append(ret, -1)
						}
					}
				}
			}
		}
	}
	// 最右元素单独处理, 0 <= i < (len-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[len(nums)-1] {
			ret = append(ret, nums[i])
			break
		}
		if i == len(nums)-2 {
			ret = append(ret, -1)
		}
	}
	return ret
}
