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
				[]int{3, 2, 2, 3},
				3,
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{0, 1, 2, 2, 3, 0, 4, 2},
				2,
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
				2,
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := removeElement(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
