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
				[]int{1, 2, 3},
			},
			[]interface{}{
				[][]int{
					{1, 2, 3},
					{1, 3, 2},
					{2, 1, 3},
					{2, 3, 1},
					{3, 1, 2},
					{3, 2, 1},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run("permute", func(t *testing.T) {
			ret := permute(c.inputs[0].([]int))
			assert.Equal(t, c.expects[0].([][]int), ret)
		})
	}
}
