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
					{15, 7},
					{9, 20},
					{3},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := levelOrderBottom(c.inputs[0].(*TreeNode))
			if !reflect.DeepEqual(ret, c.expects[0].([][]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
