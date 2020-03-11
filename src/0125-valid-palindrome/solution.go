package solution

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	r, _ := regexp.Compile("[^\\w\\d]")
	s = strings.ToLower(r.ReplaceAllString(s, ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
