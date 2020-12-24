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
				"10", "9",
			},
			[]interface{}{
				"19",
			},
		},
		{
			[]interface{}{
				"88", "933",
			},
			[]interface{}{
				"1021",
			},
		},
	}
	for _, c := range cases {
		t.Run("addStrings", func(t *testing.T) {
			ret := addStrings(c.inputs[0].(string), c.inputs[1].(string))
			assert.Equal(t, c.expects[0].(string), ret)
		})
	}
}
