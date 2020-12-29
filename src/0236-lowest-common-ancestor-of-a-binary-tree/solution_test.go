package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	p1 = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	q1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}

	q2 = &TreeNode{
		Val: 4,
	}
	p2 = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: q2,
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
				&TreeNode{
					Val:   3,
					Left:  p1,
					Right: q1,
				},
				p1,
				q1,
			},
			[]interface{}{
				&TreeNode{
					Val:   3,
					Left:  p1,
					Right: q1,
				},
			},
		},
		{
			[]interface{}{
				&TreeNode{
					Val:  3,
					Left: p2,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				p2,
				q2,
			},
			[]interface{}{
				p2,
			},
		},
	}
	for _, c := range cases {
		t.Run("lowestCommonAncestor", func(t *testing.T) {
			ret := lowestCommonAncestor(c.inputs[0].(*TreeNode), c.inputs[1].(*TreeNode), c.inputs[2].(*TreeNode))
			assert.Equal(t, c.expects[0].(*TreeNode), ret)
		})
	}
}
