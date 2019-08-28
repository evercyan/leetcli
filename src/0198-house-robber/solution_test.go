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
				[]int{1, 2, 3, 1},
			},
			[]interface{}{
				4,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 7, 9, 3, 1},
			},
			[]interface{}{
				12,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{2},
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{2, 7},
			},
			[]interface{}{
				7,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := rob(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
