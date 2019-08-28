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
				[]int{2, 1, 2},
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 2, 1},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{3, 2, 3, 4},
			},
			[]interface{}{
				10,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{3, 6, 2, 3},
			},
			[]interface{}{
				8,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := largestPerimeter(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
