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
				[]int{1, 0, -1, 0, -2, 2},
				0,
			},
			[]interface{}{
				[][]int{
					{-2, -1, 1, 2},
					{-2, 0, 0, 2},
					{-1, 0, 0, 1},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 0, -1},
				0,
			},
			[]interface{}{
				[][]int{},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{0, 0, 0, 0, 0},
				1,
			},
			[]interface{}{
				[][]int{},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := fourSum(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([][]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
