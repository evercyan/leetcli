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
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
			},
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
		},
		{
			[]interface{}{
				[]int{3, 9, 20, 15, 7},
				[]int{15, 20, 7},
			},
			[]interface{}{
				&TreeNode{
					Val: 3,
				},
			},
		},
		{
			[]interface{}{
				[]int{},
				[]int{},
			},
			[]interface{}{
				(*TreeNode)(nil),
			},
		},
	}
	for _, c := range cases {
		t.Run("buildTree", func(t *testing.T) {
			ret := buildTree(c.inputs[0].([]int), c.inputs[1].([]int))
			assert.Equal(t, c.expects[0].(*TreeNode), ret)
		})
	}
}
