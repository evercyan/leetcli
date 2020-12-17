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
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			[]interface{}{
				[]int{1, 3, 2},
			},
		},
		{
			[]interface{}{
				&TreeNode{
					Val: 1,
				},
			},
			[]interface{}{
				[]int{1},
			},
		},
	}
	for _, c := range cases {
		t.Run("inorderTraversal", func(t *testing.T) {
			ret := inorderTraversal(c.inputs[0].(*TreeNode))
			assert.Equal(t, c.expects[0].([]int), ret)
		})
	}
}
