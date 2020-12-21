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
				[]int{2, 0, 2, 1, 1, 0},
			},
			[]interface{}{
				[]int{0, 0, 1, 1, 2, 2},
			},
		},
		{
			[]interface{}{
				[]int{2, 0, 1},
			},
			[]interface{}{
				[]int{0, 1, 2},
			},
		},
	}
	for _, c := range cases {
		t.Run("sortColors", func(t *testing.T) {
			input := c.inputs[0].([]int)
			sortColors(input)
			assert.Equal(t, c.expects[0].([]int), input)
		})
	}
}
