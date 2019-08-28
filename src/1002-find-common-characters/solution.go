package solution

func commonChars(A []string) []string {
	ret := []string{}
	// 先将字符串数组解析成 char->num 数组
	sl := []map[string]int{}
	for _, s := range A {
		tmp := map[string]int{}
		for _, c := range s {
			if _, ok := tmp[string(c)]; !ok {
				tmp[string(c)] = 0
			}
			tmp[string(c)]++
		}
		sl = append(sl, tmp)
	}
	// 以第一个元素作为比较基准
	for c, n := range sl[0] {
		cmin := n
		for i := 1; i < len(sl); i++ {
			// 如果后续 map 中没有该字符, break
			if _, ok := sl[i][c]; !ok {
				cmin = 0
				break
			}
			// 比较最小出现次数
			if sl[i][c] < cmin {
				cmin = sl[i][c]
			}
		}
		if cmin <= 0 {
			continue
		}
		// 循环写入
		for i := 0; i < cmin; i++ {
			ret = append(ret, c)
		}
	}
	return ret
}
