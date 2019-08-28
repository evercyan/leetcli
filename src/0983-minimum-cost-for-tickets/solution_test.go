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
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			[]interface{}{
				11,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			[]interface{}{
				17,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
				[]int{2, 7, 15},
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mincostTickets(c.inputs[0].([]int), c.inputs[1].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
