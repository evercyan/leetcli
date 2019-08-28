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
				1,
			},
			[]interface{}{
				"1",
			},
		},
		{
			"Test 2",
			[]interface{}{
				4,
			},
			[]interface{}{
				"1211",
			},
		},
		{
			"Test 3",
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
			ret := countAndSay(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
