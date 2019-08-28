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
				121,
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				-121,
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 3",
			[]interface{}{
				10,
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isPalindrome(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
