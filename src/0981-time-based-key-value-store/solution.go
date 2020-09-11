package solution

/**
 * 常规思路, 循环去找最大值, 会超出时间限制
 * 注意: 所有 TimeMap.set 操作中的时间戳 timestamps 都是严格递增的, 直接二分法
 * Golang 的 TimeMap 结构待优化
 */

type TimeMap struct {
	Val  map[string][]string
	Time map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		Val:  map[string][]string{},
		Time: map[string][]int{},
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := tm.Val[key]; !ok {
		tm.Val[key] = []string{}
		tm.Time[key] = []int{}
	}
	tm.Val[key] = append(tm.Val[key], value)
	tm.Time[key] = append(tm.Time[key], timestamp)
	return
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.Time[key]; !ok {
		return ""
	}
	l, h := 0, len(tm.Time[key])-1
	for l <= h {
		mid := (l + h) / 2
		if tm.Time[key][mid] >= timestamp {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(tm.Time[key]) && tm.Time[key][l] == timestamp {
		return tm.Val[key][l]
	}
	if l >= 1 {
		return tm.Val[key][l-1]
	}
	return ""
}
