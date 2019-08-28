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
				[]int{1, 2, 3},
			},
			[]interface{}{
				[]int{1, 3, 2},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{3, 2, 1},
			},
			[]interface{}{
				[]int{1, 2, 3},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 1, 5},
			},
			[]interface{}{
				[]int{1, 5, 1},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nums := c.inputs[0].([]int)
			nextPermutation(nums)
			if !reflect.DeepEqual(nums, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], nums)
			}
		})
	}
}
