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
				"11",
				"1",
			},
			[]interface{}{
				"100",
			},
		},
		{
			"Test 2",
			[]interface{}{
				"1010",
				"1011",
			},
			[]interface{}{
				"10101",
			},
		},
		{
			"Test 3",
			[]interface{}{
				"111",
				"1110",
			},
			[]interface{}{
				"10101",
			},
		},
		{
			"Test 4",
			[]interface{}{
				"0",
				"0",
			},
			[]interface{}{
				"0",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := addBinary(c.inputs[0].(string), c.inputs[1].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
