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
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			[]interface{}{
				[][]int{
					{3},
					{9, 20},
					{15, 7},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("levelOrder", func(t *testing.T) {
			ret := levelOrder(c.inputs[0].(*TreeNode))
			assert.Equal(t, c.expects[0].([][]int), ret)
		})
	}
}
