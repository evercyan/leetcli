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
			"Test",
			[]interface{}{
				"aa",
				"a",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				"aa",
				"*",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				"cb",
				"?a",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				"adceb",
				"*a*b",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				"acdcb",
				"a*c?b",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				"",
				"",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				"",
				"a",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				"a",
				"",
			},
			[]interface{}{
				false,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.inputs[0].(string), c.inputs[1].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
