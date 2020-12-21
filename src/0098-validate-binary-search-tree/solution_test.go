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
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			[]interface{}{
				true,
			},
		},
		{
			[]interface{}{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run("isValidBST", func(t *testing.T) {
			ret := isValidBST(c.inputs[0].(*TreeNode))
			assert.Equal(t, c.expects[0].(bool), ret)
		})
	}
}
