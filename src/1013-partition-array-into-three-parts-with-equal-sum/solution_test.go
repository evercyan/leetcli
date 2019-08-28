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
				[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{3, 3},
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 5",
			[]interface{}{
				[]int{3, 4, 7, 2},
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canThreePartsEqualSum(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
