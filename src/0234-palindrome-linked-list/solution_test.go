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
					},
				},
			},
			[]interface{}{
				false,
			},
		},
		{
			[]interface{}{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
			[]interface{}{
				true,
			},
		},
	}
	for _, c := range cases {
		t.Run("isPalindrome", func(t *testing.T) {
			ret := isPalindrome(c.inputs[0].(*ListNode))
			assert.Equal(t, c.expects[0].(bool), ret)
		})
	}
}
