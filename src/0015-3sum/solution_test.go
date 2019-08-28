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
				[]int{-1, 0, 1, 2, -1, -4},
			},
			[]interface{}{
				[][]int{
					[]int{-1, -1, 2},
					[]int{-1, 0, 1},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{-1, 0},
			},
			[]interface{}{
				[][]int{},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{0, 1, 1, 1, 0, 0, -1, -1, -1},
			},
			[]interface{}{
				[][]int{
					[]int{-1, 0, 1},
					[]int{0, 0, 0},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := threeSum(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].([][]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
