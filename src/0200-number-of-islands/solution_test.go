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
					[]byte("11110"),
					[]byte("11010"),
					[]byte("11000"),
					[]byte("00000"),
				},
			},
			[]interface{}{
				1,
			},
		},
		{
			[]interface{}{
				[][]byte{
					[]byte("11000"),
					[]byte("11000"),
					[]byte("00100"),
					[]byte("00011"),
				},
			},
			[]interface{}{
				3,
			},
		},
	}
	for _, c := range cases {
		t.Run("numIslands", func(t *testing.T) {
			ret := numIslands(c.inputs[0].([][]byte))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
