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
				"2",
				"3",
			},
			[]interface{}{
				"6",
			},
		},
		{
			"Test 2",
			[]interface{}{
				"12",
				"34",
			},
			[]interface{}{
				"408",
			},
		},
		{
			"Test 3",
			[]interface{}{
				"123",
				"456",
			},
			[]interface{}{
				"56088",
			},
		},
		{
			"Test 4",
			[]interface{}{
				"0",
				"456",
			},
			[]interface{}{
				"0",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := multiply(c.inputs[0].(string), c.inputs[1].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
