package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func int2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func int2ListVsCycle(nums []int, pos int) *ListNode {
	head := int2List(nums)
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

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				int2ListVsCycle([]int{3, 2, 0, -4}, 1),
			},
			[]interface{}{
				// 此处非简单的 int2List([]int{2, 0, -4}), 而一个环形链表
				// pos 为 1, 则 expect 为预期的 x.Next, 为 2 则为 x.Next.Next
				int2ListVsCycle([]int{3, 2, 0, -4}, 1).Next,
			},
		},
		{
			[]interface{}{
				int2ListVsCycle([]int{3, 2, 0, -4}, -1),
			},
			[]interface{}{
				(*ListNode)(nil),
			},
		},
	}

	for _, c := range cases {
		t.Run("detectCycle", func(t *testing.T) {
			ret := detectCycle(c.inputs[0].(*ListNode))
			assert.Equal(t, c.expects[0].(*ListNode), ret)
		})
	}
}
