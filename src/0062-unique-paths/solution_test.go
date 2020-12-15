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
				3, 7,
			},
			[]interface{}{
				28,
			},
		},
		{
			[]interface{}{
				3, 2,
			},
			[]interface{}{
				3,
			},
		},
	}
	for _, c := range cases {
		t.Run("uniquePaths", func(t *testing.T) {
			ret := uniquePaths(c.inputs[0].(int), c.inputs[1].(int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
