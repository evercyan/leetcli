package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"ABCCED",
			},
			[]interface{}{
				true,
			},
		},
		{
			[]interface{}{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"SEE",
			},
			[]interface{}{
				true,
			},
		},
		{
			[]interface{}{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"ABCB",
			},
			[]interface{}{
				false,
			},
		},
	}
	for _, c := range cases {
		t.Run("exist", func(t *testing.T) {
			ret := exist(c.inputs[0].([][]byte), c.inputs[1].(string))
			assert.Equal(t, c.expects[0].(bool), ret)
		})
	}
}
