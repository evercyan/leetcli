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
				2,
			},
			[]interface{}{
				"110",
			},
		},
		{
			"Test 2",
			[]interface{}{
				3,
			},
			[]interface{}{
				"111",
			},
		},
		{
			"Test 3",
			[]interface{}{
				4,
			},
			[]interface{}{
				"100",
			},
		},
		{
			"Test 4",
			[]interface{}{
				0,
			},
			[]interface{}{
				"0",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := baseNeg2(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
