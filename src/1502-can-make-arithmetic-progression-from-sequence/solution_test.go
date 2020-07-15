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
				[]int{3, 5, 1},
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]int{1, 2, 4},
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]int{1, 2},
			},
			[]interface{}{
				true,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canMakeArithmeticProgression(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
