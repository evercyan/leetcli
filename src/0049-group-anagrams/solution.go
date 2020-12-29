package solution

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		// 对 str 排序, 使字母异位词得到统一的 key 值
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		key := string(b)
		m[key] = append(m[key], str)
	}

	var ret [][]string
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}
