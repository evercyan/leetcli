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
				[]int{1, 2, 3, 0, 0, 0},
				3,
				[]int{2, 5, 6},
				3,
			},
			[]interface{}{
				[]int{1, 2, 2, 3, 5, 6},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 5, 6, 0, 0, 0},
				3,
				[]int{1, 2, 3},
				3,
			},
			[]interface{}{
				[]int{1, 2, 2, 3, 5, 6},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			merge(c.inputs[0].([]int), c.inputs[1].(int), c.inputs[2].([]int), c.inputs[3].(int))
			if !reflect.DeepEqual(c.inputs[0].([]int), c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], c.inputs[0].([]int))
			}
		})
	}
}
