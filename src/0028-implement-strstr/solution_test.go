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
				"hello",
				"ll",
			},
			[]interface{}{
				2,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"aaaaa",
				"bba",
			},
			[]interface{}{
				-1,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"",
				"bba",
			},
			[]interface{}{
				-1,
			},
		},
		{
			"Test 4",
			[]interface{}{
				"aaaaa",
				"",
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 5",
			[]interface{}{
				"mississippi",
				"issip",
			},
			[]interface{}{
				4,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := strStr(c.inputs[0].(string), c.inputs[1].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
