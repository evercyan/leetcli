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
				[]int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			[]interface{}{
				[]int{5, 6},
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				[]int{},
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findDisappearedNumbers(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
