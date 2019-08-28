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
				[]int{1, 4},
				[]int{2},
			},
			[]interface{}{
				2.0,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 4},
				[]int{2, 3},
			},
			[]interface{}{
				2.5,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findMedianSortedArrays(c.inputs[0].([]int), c.inputs[1].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(float64)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
