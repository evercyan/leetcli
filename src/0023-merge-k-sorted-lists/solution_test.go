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
				[]*ListNode{
					&ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
					&ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
					&ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 5,
											Next: &ListNode{
												Val: 6,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]*ListNode{
					&ListNode{
						Val: 1,
					},
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
				[]*ListNode{},
			},
			[]interface{}{
				nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mergeKLists(c.inputs[0].([]*ListNode))
			if c.expects[0] == nil {
				// nil 特殊处理
				if ret != nil {
					t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
				}
			} else {
				if !reflect.DeepEqual(ret, c.expects[0].(*ListNode)) {
					t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
				}
			}
		})
	}
}
