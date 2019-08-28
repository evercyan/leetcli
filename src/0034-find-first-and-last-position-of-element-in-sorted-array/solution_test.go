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
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			[]interface{}{
				[]int{3, 4},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{5, 7, 7, 8, 8, 10},
				6,
			},
			[]interface{}{
				[]int{-1, -1},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
				3,
			},
			[]interface{}{
				[]int{-1, -1},
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{3, 3, 3},
				3,
			},
			[]interface{}{
				[]int{0, 2},
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{1, 4},
				4,
			},
			[]interface{}{
				[]int{1, 1},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchRange(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
