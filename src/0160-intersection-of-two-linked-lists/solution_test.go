package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// line 33, 42, 48 处需用此变量
	l1 = &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	l2 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
)

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				&ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
						Next: l1,
					},
				},
				&ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: l1,
						},
					},
				},
			},
			[]interface{}{
				l1,
			},
		},
		{
			[]interface{}{
				&ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  1,
							Next: l2,
						},
					},
				},
				&ListNode{
					Val:  3,
					Next: l2,
				},
			},
			[]interface{}{
				l2,
			},
		},
		{
			[]interface{}{
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
			[]interface{}{
				(*ListNode)(nil),
			},
		},
	}

	for _, c := range cases {
		t.Run("getIntersectionNode", func(t *testing.T) {
			ret := getIntersectionNode(c.inputs[0].(*ListNode), c.inputs[1].(*ListNode))
			assert.Equal(t, c.expects[0].(*ListNode), ret)
		})
	}
}
