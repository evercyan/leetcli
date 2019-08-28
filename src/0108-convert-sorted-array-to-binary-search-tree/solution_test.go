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
				[]int{-10, -3, 0, 5, 9},
			},
			[]interface{}{
				&TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val: -10,
						},
					},
					Right: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := sortedArrayToBST(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(*TreeNode)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
