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
				[]int{10, 15, 20},
			},
			[]interface{}{
				15,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			[]interface{}{
				6,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1},
			},
			[]interface{}{
				0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := minCostClimbingStairs(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
