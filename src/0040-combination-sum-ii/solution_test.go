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
				[]int{10, 1, 2, 7, 6, 1, 5},
				8,
			},
			[]interface{}{
				[][]int{
					[]int{1, 1, 6},
					[]int{1, 2, 5},
					[]int{1, 7},
					[]int{2, 6},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 5, 2, 1, 2},
				5,
			},
			[]interface{}{
				[][]int{
					[]int{1, 2, 2},
					[]int{5},
				},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{},
				8,
			},
			[]interface{}{
				[][]int{},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := combinationSum2(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([][]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
