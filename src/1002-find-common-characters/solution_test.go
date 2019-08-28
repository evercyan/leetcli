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
				[]string{"bella", "label", "roller"},
			},
			[]interface{}{
				[]string{"e", "l", "l"},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]string{"cool", "lock", "cook"},
			},
			[]interface{}{
				[]string{"c", "o"},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := commonChars(c.inputs[0].([]string))
			if !reflect.DeepEqual(len(ret), len(c.expects[0].([]string))) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
