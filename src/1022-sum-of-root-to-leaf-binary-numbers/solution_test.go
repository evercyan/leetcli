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
					Val: 1,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			[]interface{}{
				22,
			},
		},
		{
			"Test 2",
			[]interface{}{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			[]interface{}{
				4,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sumRootToLeaf(c.inputs[0].(*TreeNode))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
