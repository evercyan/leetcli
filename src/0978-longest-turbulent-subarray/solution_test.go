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
				[]int{9, 4, 2, 10, 7, 8, 9, 1, 9},
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{9},
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0},
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{0, 1, 0, 0},
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 6",
			[]interface{}{
				[]int{4, 8, 12, 16},
			},
			[]interface{}{
				2,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxTurbulenceSize(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
