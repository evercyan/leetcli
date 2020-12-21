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
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[]interface{}{
				[][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("rotate", func(t *testing.T) {
			input := c.inputs[0].([][]int)
			rotate(input)
			assert.Equal(t, c.expects[0].([][]int), input)
		})
	}
}
