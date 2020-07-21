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
			"Test",
			[]interface{}{
				5,
				[]int{},
				[]int{},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test",
			[]interface{}{
				5,
				[]int{1, 1, 1},
				[]int{1, 1, 1},
			},
			[]interface{}{
				3,
			},
		},
		{
			"Test",
			[]interface{}{
				4,
				[]int{1, 4, 3},
				[]int{15, 30, 20},
			},
			[]interface{}{
				35,
			},
		},
		{
			"Test",
			[]interface{}{
				15,
				[]int{5, 4, 7, 2, 6},
				[]int{12, 3, 10, 3, 6},
			},
			[]interface{}{
				25,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := dpPackage(c.inputs[0].(int), c.inputs[1].([]int), c.inputs[2].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
