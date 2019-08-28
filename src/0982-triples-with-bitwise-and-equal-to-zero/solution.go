package solution

/**
 * 题解
 *
 * 先双层循环, 取 i & j 结果, 存 ijMap, 结果值累加
 *  单循环 k, 再循环 ijMap, k & mk, 次数 mv 累加
 */

func countTriplets(A []int) int {
	if len(A) <= 0 {
		return 0
	}
	ret := 0
	ijMap := map[int]int{}
	for _, i := range A {
		for _, j := range A {
			if _, ok := ijMap[i&j]; !ok {
				ijMap[i&j] = 0
			}
			ijMap[i&j]++
		}
	}
	for _, k := range A {
		for mk, mv := range ijMap {
			if k&mk == 0 {
				ret += mv
			}
		}
	}
	return ret
}
