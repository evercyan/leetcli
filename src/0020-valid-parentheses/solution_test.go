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
				"()",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"()[]{}",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"(]",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 4",
			[]interface{}{
				"([)]",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 5",
			[]interface{}{
				"{[]}",
			},
			[]interface{}{
				true,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValid(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
