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
				5,
			},
			[]interface{}{
				[][]int{
					{1},
					{1, 1},
					{1, 2, 1},
					{1, 3, 3, 1},
					{1, 4, 6, 4, 1},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("generate", func(t *testing.T) {
			ret := generate(c.inputs[0].(int))
			assert.Equal(t, c.expects[0].([][]int), ret)
		})
	}
}
