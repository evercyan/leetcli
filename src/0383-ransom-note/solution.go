package solution

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for i := 0; i < len(ransomNote); i++ {
		char := string(ransomNote[i])
		if !strings.Contains(magazine, char) {
			return false
		}
		magazine = strings.Replace(magazine, char, "", 1)
	}
	return true
}
