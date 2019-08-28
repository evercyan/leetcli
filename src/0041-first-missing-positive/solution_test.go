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
				[]int{1, 2, 0},
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{3, 4, -1, 1},
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{7, 8, 9, 11, 12},
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{1},
			},
			[]interface{}{
				2,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := firstMissingPositive(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
