package solution

/**
 * 直觉是维护 min, 但每次 pop 后, 需遍历计算, O(n)
 * 用空间换时间
 * push(x) 前, 先 push(min), pop 的时候, pop() 两次, pop 完取 min 直接取 list[1]
 */

type MinStack struct {
	min  int
	list []int
}

func Constructor() MinStack {
	return MinStack{
		list: []int{},
	}
}

func (s *MinStack) Push(x int) {
	var min int
	if len(s.list) <= 0 {
		min = x
	} else {
		if x < s.list[1] {
			min = x
		} else {
			min = s.list[1]
		}
	}
	s.list = append([]int{x, min}, s.list...)
}

func (s *MinStack) Pop() {
	if len(s.list) < 2 {
		return
	}
	s.list = append([]int{}, s.list[2:]...)
}

func (s *MinStack) Top() int {
	if len(s.list) <= 0 {
		return 0
	}
	return s.list[0]
}

func (s *MinStack) GetMin() int {
	if len(s.list) <= 0 {
		return 0
	}
	return s.list[1]
}
