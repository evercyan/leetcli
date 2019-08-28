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

/**
 * 队列初始化
 */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		len:   k,
		data:  make([]int, k),
		front: 0,
		rear:  0,
	}
}

/**
 * 尾部插入元素
 */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		// 队列已满
		return false
	}
	this.data[this.rear] = value
	// 防止越界
	this.rear = (this.rear + 1) % this.len
	return true
}

/**
 * 头部删除元素
 */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		// 队列为空
		return false
	}
	// 清空数据
	this.data[this.front] = 0
	this.front = (this.front + 1) % this.len
	return true
}

/**
 * 获取队列头部的元素
 */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/**
 * 获取队列尾部的元素
 */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.rear == 0 {
		return this.data[this.len-1]
	}
	return this.data[this.rear-1]
}

/**
 * 判断队列为空的条件
 * rear == front 且 front 指向为 0
 */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear && this.data[this.front] == 0
}

/**
 * 判断队列为满的条件
 * rear == front 且 front 指向不为 0
 */
func (this *MyCircularQueue) IsFull() bool {
	return this.front == this.rear && this.data[this.front] != 0
}
