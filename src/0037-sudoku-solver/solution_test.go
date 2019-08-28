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
				[][]byte{
					[]byte("..9748..."),
					[]byte("7........"),
					[]byte(".2.1.9..."),
					[]byte("..7...24."),
					[]byte(".64.1.59."),
					[]byte(".98...3.."),
					[]byte("...8.3.2."),
					[]byte("........6"),
					[]byte("...2759.."),
				},
			},
			[]interface{}{
				[][]byte{
					[]byte("519748632"),
					[]byte("783652419"),
					[]byte("426139875"),
					[]byte("357986241"),
					[]byte("264317598"),
					[]byte("198524367"),
					[]byte("975863124"),
					[]byte("832491756"),
					[]byte("641275983"),
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			solveSudoku(c.inputs[0].([][]byte))
			if !reflect.DeepEqual(c.inputs[0], c.expects[0].([][]byte)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], c.inputs[0])
			}
		})
	}
}
