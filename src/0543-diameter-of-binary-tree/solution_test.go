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
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			[]interface{}{
				3,
			},
		},
		{
			[]interface{}{
				&TreeNode{},
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run("diameterOfBinaryTree", func(t *testing.T) {
			ret := diameterOfBinaryTree(c.inputs[0].(*TreeNode))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
