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
				[]int{1, 2, 1},
			},
			[]interface{}{
				[]int{2, -1, 2},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 1, 1},
			},
			[]interface{}{
				[]int{-1, 2, 2},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 1, 2},
			},
			[]interface{}{
				[]int{2, 2, -1},
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				[]int{},
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{1},
			},
			[]interface{}{
				[]int{-1},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := nextGreaterElements(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
