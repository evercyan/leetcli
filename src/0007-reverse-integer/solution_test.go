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
				123,
			},
			[]interface{}{
				321,
			},
		},
		{
			"Test 2",
			[]interface{}{
				-123,
			},
			[]interface{}{
				-321,
			},
		},
		{
			"Test 3",
			[]interface{}{
				120,
			},
			[]interface{}{
				21,
			},
		},
		{
			"Test 4",
			[]interface{}{
				10000000001,
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverse(c.inputs[0].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
