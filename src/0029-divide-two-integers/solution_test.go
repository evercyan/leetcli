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
				10,
				3,
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 2",
			[]interface{}{
				7,
				-3,
			},
			[]interface{}{
				-2,
			},
		},
		{
			"Test 3",
			[]interface{}{
				0,
				3,
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 4",
			[]interface{}{
				-2147483648,
				-1,
			},
			[]interface{}{
				2147483647,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divide(c.inputs[0].(int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
