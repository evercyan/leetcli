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
				5,
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 2",
			[]interface{}{
				7,
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 3",
			[]interface{}{
				10,
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 4",
			[]interface{}{
				0,
			},
			[]interface{}{
				1,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := bitwiseComplement(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
