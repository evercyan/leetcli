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
				[]int{1, 3, 5, 6},
				5,
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 3, 5, 6},
				2,
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 3, 5, 6},
				7,
			},
			[]interface{}{
				4,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{1, 3, 5, 6},
				0,
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{},
				0,
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchInsert(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
