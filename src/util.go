package src

// src 下各题目解答均为独立 package, 通用的代码封装于此, 方便复用

// -------- Struct

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// -------- ListNode

// list2node []int 转 *ListNode
func list2node(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// node2list *ListNode 转 []int
func node2list(l *ListNode) []int {
	var r []int
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}

// list2NodeVsCycle 构造环型链表
func list2NodeVsCycle(nums []int, pos int) *ListNode {
	head := list2node(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

// -------- Math

// min ...
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// max ...
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// abs ...
func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// -------- transfer

// s2i string 2 int
func s2i(v string) int {
	r, _ := strconv.Atoi(v)
	return r
}

// i2s int 2 string
func i2s(v int) string {
	return fmt.Sprintf("%v", v)
}
