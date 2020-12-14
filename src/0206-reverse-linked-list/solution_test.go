package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			[]interface{}{
				&ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 1,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("reverseList", func(t *testing.T) {
			ret := reverseList(c.inputs[0].(*ListNode))
			assert.Equal(t, c.expects[0].(*ListNode), ret)
			//assert.True(t, reflect.DeepEqual(ret, c.expects[0].(*ListNode)))
		})
	}
}
