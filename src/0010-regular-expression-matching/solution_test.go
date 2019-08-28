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
				"aa",
				"a",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"aa",
				"a*",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"ab",
				".*",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 4",
			[]interface{}{
				"aab",
				"c*a*b",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 5",
			[]interface{}{
				"mississippi",
				"mis*is*p*.",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 6",
			[]interface{}{
				"",
				"",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 7",
			[]interface{}{
				"",
				"a",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 8",
			[]interface{}{
				"a",
				"",
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.inputs[0].(string), c.inputs[1].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
