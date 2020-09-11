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
				[]int{1, 2, 2, 6, 6, 6, 6, 7, 10},
			},
			[]interface{}{
				6,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{6},
			},
			[]interface{}{
				6,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 2, 3, 4, 5},
			},
			[]interface{}{
				5,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[]int{1, 1, 2, 2, 3, 3, 3, 3},
			},
			[]interface{}{
				3,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSpecialInteger(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
