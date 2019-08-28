package solution

import (
	"reflect"
	"testing"
)

func node2list(l *ListNode) []int {
	r := []int{}
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}

func list2node(l []int) *ListNode {
	r := &ListNode{}
	t := r
	for _, v := range l {
		t.Val = v
		t.Next = &ListNode{}
		t = t.Next
	}
	return r
}

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test 1",
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			[]interface{}{
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				&ListNode{
					Val: 1,
				},
			},
			[]interface{}{
				&ListNode{
					Val: 1,
				},
			},
		},
		{
			"Test 3",
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			[]interface{}{
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := swapPairs(c.inputs[0].(*ListNode))
			if !reflect.DeepEqual(ret, c.expects[0].(*ListNode)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, node2list(c.expects[0].(*ListNode)), node2list(ret))
			}
		})
	}
}
