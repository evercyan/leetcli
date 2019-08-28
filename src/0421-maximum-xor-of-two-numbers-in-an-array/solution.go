package solution

func findMaximumXOR(nums []int) int {
	max, mask := 0, 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << uint(i))
		set := make(map[int]bool)
		arr := make([]int, 0, len(nums))
		for _, v := range nums {
			t := v & mask
			if _, ok := set[t]; !ok {
				set[t] = true
				arr = append(arr, t)
			}
		}
		tmp := max | (1 << uint(i))
		for _, pre := range arr {
			if _, ok := set[tmp^pre]; ok {
				max = tmp
				break
			}
		}
	}
	return max
}
