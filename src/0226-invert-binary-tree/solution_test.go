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
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			[]interface{}{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("invertTree", func(t *testing.T) {
			ret := invertTree(c.inputs[0].(*TreeNode))
			assert.Equal(t, c.expects[0].(*TreeNode), ret)
		})
	}
}
