package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test 1",
			[]interface{}{
				&ListNode{Val: 1},
				&ListNode{Val: 9, Next: &ListNode{Val: 9}},
			},
			[]interface{}{
				&ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
			},
		},
		{
			"Test 2",
			[]interface{}{
				&ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}},
				&ListNode{Val: 9, Next: &ListNode{Val: 8, Next: &ListNode{Val: 7}}},
			},
			[]interface{}{
				&ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1}}}},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addTwoNumbers(c.inputs[0].(*ListNode), c.inputs[1].(*ListNode))
			if !reflect.DeepEqual(ret, c.expects[0].(*ListNode)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
