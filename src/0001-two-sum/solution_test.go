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
				[]int{2, 7, 11},
				9,
			},
			[]interface{}{
				[]int{0, 1},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 3, 4},
				7,
			},

			[]interface{}{
				[]int{1, 2},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{7, 6, 5, 3, 2, 1, 4, 9, 10},
				17,
			},
			[]interface{}{
				[]int{0, 8},
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{1, 2},
				17,
			},
			[]interface{}{
				[]int{},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoSum(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
