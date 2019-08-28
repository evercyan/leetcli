package solution

type Interval struct {
	Start int
	End   int
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
func insert(intervals []Interval, newInterval Interval) []Interval {
	var r []Interval
	iLen := len(intervals)
	if iLen <= 0 {
		r = append(r, newInterval)
		return r
	}
	var left []Interval
	var right []Interval
	s := newInterval.Start
	e := newInterval.End
	for i := 0; i < iLen; i++ {
		if intervals[i].End < s {
			left = append(left, intervals[i])
		} else if intervals[i].Start > e {
			right = append(right, intervals[i])
		} else {
			s = min(s, intervals[i].Start)
			e = max(e, intervals[i].End)
		}
	}
	mid := &Interval{
		Start: s,
		End:   e,
	}
	r = append(r, left...)
	r = append(r, *mid)
	r = append(r, right...)
	return r
}
