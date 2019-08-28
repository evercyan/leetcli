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
				"abcabc",
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"hhhhh",
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"pwwkew",
			},
			[]interface{}{
				3,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lengthOfLongestSubstring(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
