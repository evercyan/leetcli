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
				"42",
			},
			[]interface{}{
				42,
			},
		},
		{
			"Test 2",
			[]interface{}{
				"   -42",
			},
			[]interface{}{
				-42,
			},
		},
		{
			"Test 3",
			[]interface{}{
				"4193 with words",
			},
			[]interface{}{
				4193,
			},
		},
		{
			"Test 4",
			[]interface{}{
				"words and 987",
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 5",
			[]interface{}{
				"-91283472332",
			},
			[]interface{}{
				-2147483648,
			},
		},
		{
			"Test 6",
			[]interface{}{
				"",
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 7",
			[]interface{}{
				"+10",
			},
			[]interface{}{
				10,
			},
		},
		{
			"Test 8",
			[]interface{}{
				"91283472339",
			},
			[]interface{}{
				2147483647,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := myAtoi(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
