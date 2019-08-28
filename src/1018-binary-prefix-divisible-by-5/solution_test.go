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
				[]int{0, 1, 1},
			},
			[]interface{}{
				[]bool{true, false, false},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 1, 1},
			},
			[]interface{}{
				[]bool{false, false, false},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{0, 1, 1, 1, 1, 1},
			},
			[]interface{}{
				[]bool{true, false, false, false, true, false},
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{1, 1, 1, 0, 1},
			},
			[]interface{}{
				[]bool{false, false, false, false, false},
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				[]bool{},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := prefixesDivBy5(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].([]bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
