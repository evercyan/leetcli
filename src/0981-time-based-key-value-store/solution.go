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

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.Val[key]; !ok {
		this.Val[key] = []string{}
		this.Time[key] = []int{}
	}
	this.Val[key] = append(this.Val[key], value)
	this.Time[key] = append(this.Time[key], timestamp)
	return
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.Time[key]; !ok {
		return ""
	}
	l, h := 0, len(this.Time[key])-1
	for l <= h {
		mid := (l + h) / 2
		if this.Time[key][mid] >= timestamp {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(this.Time[key]) && this.Time[key][l] == timestamp {
		return this.Val[key][l]
	}
	if l >= 1 {
		return this.Val[key][l-1]
	}
	return ""
}
