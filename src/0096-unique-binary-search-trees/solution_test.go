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
				3,
			},
			[]interface{}{
				5,
			},
		},
	}
	for _, c := range cases {
		t.Run("numTrees", func(t *testing.T) {
			ret := numTrees(c.inputs[0].(int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
