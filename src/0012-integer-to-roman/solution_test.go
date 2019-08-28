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
				3,
			},
			[]interface{}{
				"III",
			},
		},
		{
			"Test 2",
			[]interface{}{
				4,
			},
			[]interface{}{
				"IV",
			},
		},
		{
			"Test 3",
			[]interface{}{
				9,
			},
			[]interface{}{
				"IX",
			},
		},
		{
			"Test 4",
			[]interface{}{
				58,
			},
			[]interface{}{
				"LVIII",
			},
		},
		{
			"Test 5",
			[]interface{}{
				1994,
			},
			[]interface{}{
				"MCMXCIV",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := intToRoman(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
