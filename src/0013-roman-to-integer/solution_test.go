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
				"III",
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"IV",
			},
			[]interface{}{
				4,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"IX",
			},
			[]interface{}{
				9,
			},
		},
		{
			"Test 4",
			[]interface{}{
				"LVIII",
			},
			[]interface{}{
				58,
			},
		},
		{
			"Test 5",
			[]interface{}{
				"MCMXCIV",
			},
			[]interface{}{
				1994,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := romanToInt(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
