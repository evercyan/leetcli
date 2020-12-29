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
				[]int{1, 3, 3, 1},
			},
		},
	}
	for _, c := range cases {
		t.Run("getRow", func(t *testing.T) {
			ret := getRow(c.inputs[0].(int))
			assert.Equal(t, c.expects[0].([]int), ret)
		})
	}
}
