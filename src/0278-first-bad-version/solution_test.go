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
				4,
			},
		},
	}
	for _, c := range cases {
		t.Run("isBadVersion", func(t *testing.T) {
			ret := firstBadVersion(c.inputs[0].(int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
