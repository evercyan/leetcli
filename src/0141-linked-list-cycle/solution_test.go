package solution

import (
	"reflect"
	"testing"
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
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test",
			[]interface{}{
				int2ListVsCycle([]int{3, 2, 0, -4}, 1),
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				int2ListVsCycle([]int{1, 2}, 0),
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				int2ListVsCycle([]int{1, 2}, -1),
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				&ListNode{},
			},
			[]interface{}{
				false,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := hasCycle(c.inputs[0].(*ListNode))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
			ret1 := hasCycleByHash(c.inputs[0].(*ListNode))
			if !reflect.DeepEqual(ret1, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret1)
			}
		})
		index++
	}
}
