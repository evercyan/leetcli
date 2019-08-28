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
				[]int{2, 3, 6, 7},
				7,
			},
			[]interface{}{
				[][]int{
					[]int{2, 2, 3},
					[]int{7},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{2, 3, 5},
				8,
			},
			[]interface{}{
				[][]int{
					[]int{2, 2, 2, 2},
					[]int{2, 3, 3},
					[]int{3, 5},
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
			ret := combinationSum(c.inputs[0].([]int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([][]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
