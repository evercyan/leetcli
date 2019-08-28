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
				"(()())(())",
			},
			[]interface{}{
				"()()()",
			},
		},
		{
			"Test 2",
			[]interface{}{
				"(()())(())(()(()))",
			},
			[]interface{}{
				"()()()()(())",
			},
		},
		{
			"Test 3",
			[]interface{}{
				"()()",
			},
			[]interface{}{
				"",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := removeOuterParentheses(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
