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
				"LEETCODEISHIRING",
				3,
			},
			[]interface{}{
				"LCIRETOESIIGEDHN",
			},
		},
		{
			"Test 2",
			[]interface{}{
				"LEETCODEISHIRING",
				4,
			},
			[]interface{}{
				"LDREOEIIECIHNTSG",
			},
		},
		{
			"Test 2",
			[]interface{}{
				"LEETCODEISHIRING",
				1,
			},
			[]interface{}{
				"LEETCODEISHIRING",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convert(c.inputs[0].(string), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
