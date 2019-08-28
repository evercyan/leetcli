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
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
			[]interface{}{
				[]int{5, 5, 0},
			},
		},
		{
			"Test 2",
			[]interface{}{
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			[]interface{}{
				[]int{7, 0, 5, 5, 0},
			},
		},
		{
			"Test 3",
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val: 5,
											Next: &ListNode{
												Val: 1,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			[]interface{}{
				[]int{7, 9, 9, 9, 0, 5, 0, 0},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := nextLargerNodes(c.inputs[0].(*ListNode))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
