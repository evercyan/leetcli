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
				"A man, a plan, a canal: Panama",
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test",
			[]interface{}{
				"race a car",
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test",
			[]interface{}{
				"",
			},
			[]interface{}{
				true,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isPalindrome(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
