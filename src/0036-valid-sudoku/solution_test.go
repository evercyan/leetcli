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
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			[]interface{}{
				true,
			},
		},
		{
			"Test 2",
			[]interface{}{
				[][]byte{
					[]byte("83..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 3",
			[]interface{}{
				[][]byte{
					[]byte("53..5...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			[]interface{}{
				false,
			},
		},
		{
			"Test 4",
			[]interface{}{
				[][]byte{
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".93....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValidSudoku(c.inputs[0].([][]byte))
			if !reflect.DeepEqual(ret, c.expects[0].(bool)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
