package solution

import (
	"strings"
)

func isValid(s string) bool {
	q := s
	for i := 0; i < len(s); i++ {
		if q == "" {
			break
		}
		q = strings.Replace(q, "{}", "", -1)
		q = strings.Replace(q, "[]", "", -1)
		q = strings.Replace(q, "()", "", -1)
	}
	return q == ""
}
