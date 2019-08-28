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
				[][]int{
					[]int{0, 0, 0, 0},
					[]int{1, 0, 1, 0},
					[]int{0, 1, 1, 0},
					[]int{0, 0, 0, 0},
				},
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[][]int{
					[]int{0, 1, 1, 0},
					[]int{0, 0, 1, 0},
					[]int{0, 0, 1, 0},
					[]int{0, 0, 0, 0},
				},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[][]int{
					[]int{0, 1, 1, 0},
					[]int{0, 0, 1, 1},
					[]int{0, 0, 1, 0},
					[]int{0, 0, 0, 0},
				},
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := numEnclaves(c.inputs[0].([][]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
