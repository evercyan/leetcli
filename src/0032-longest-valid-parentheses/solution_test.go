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
				"(()",
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 2",
			[]interface{}{
				")()())",
			},
			[]interface{}{
				4,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestValidParentheses(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
