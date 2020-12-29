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
				"a", "b",
			},
			[]interface{}{
				false,
			},
		},
		{
			[]interface{}{
				"aa", "ab",
			},
			[]interface{}{
				false,
			},
		},
		{
			[]interface{}{
				"aa", "aab",
			},
			[]interface{}{
				true,
			},
		},
	}
	for _, c := range cases {
		t.Run("canConstruct", func(t *testing.T) {
			ret := canConstruct(c.inputs[0].(string), c.inputs[1].(string))
			assert.Equal(t, c.expects[0].(bool), ret)
		})
	}
}
