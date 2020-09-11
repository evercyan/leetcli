package solution

/**
 * 循环队列
 */

type MyCircularQueue struct {
	data  []int
	len   int
	front int
	rear  int
}

// Constructor 队列初始化
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		len:   k,
		data:  make([]int, k),
		front: 0,
		rear:  0,
	}
}

// EnQueue 尾部插入元素
func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		// 队列已满
		return false
	}
	cq.data[cq.rear] = value
	// 防止越界
	cq.rear = (cq.rear + 1) % cq.len
	return true
}

// DeQueue 头部删除元素
func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		// 队列为空
		return false
	}
	// 清空数据
	cq.data[cq.front] = 0
	cq.front = (cq.front + 1) % cq.len
	return true
}

// Front 获取队列头部的元素
func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.data[cq.front]
}

// Rear 获取队列尾部的元素
func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	if cq.rear == 0 {
		return cq.data[cq.len-1]
	}
	return cq.data[cq.rear-1]
}

// IsEmpty 判断队列为空的条件 rear == front 且 front 指向为 0
func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.front == cq.rear && cq.data[cq.front] == 0
}

// IsFull 判断队列为满的条件 rear == front 且 front 指向不为 0
func (cq *MyCircularQueue) IsFull() bool {
	return cq.front == cq.rear && cq.data[cq.front] != 0
}
