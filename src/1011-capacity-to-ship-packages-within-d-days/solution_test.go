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
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				5,
			},
			[]interface{}{
				15,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{3, 2, 2, 4, 1, 4},
				3,
			},
			[]interface{}{
				6,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 2, 3, 1, 1},
				4,
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{3, 9, 6, 2, 5, 12},
				4,
			},
			[]interface{}{
				12,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := shipWithinDays(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
