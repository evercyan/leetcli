package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test 1",
			[]interface{}{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				22,
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				1,
			},
			[]interface{}{
				false,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := hasPathSum(c.inputs[0].(*TreeNode), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
