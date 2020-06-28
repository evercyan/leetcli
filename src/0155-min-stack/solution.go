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

func (this *MinStack) Push(x int) {
	var min int
	if len(this.list) <= 0 {
		min = x
	} else {
		if x < this.list[1] {
			min = x
		} else {
			min = this.list[1]
		}
	}
	this.list = append([]int{x, min}, this.list...)
}

func (this *MinStack) Pop() {
	if len(this.list) < 2 {
		return
	}
	this.list = append([]int{}, this.list[2:]...)
}

func (this *MinStack) Top() int {
	if len(this.list) <= 0 {
		return 0
	}
	return this.list[0]
}

func (this *MinStack) GetMin() int {
	if len(this.list) <= 0 {
		return 0
	}
	return this.list[1]
}
