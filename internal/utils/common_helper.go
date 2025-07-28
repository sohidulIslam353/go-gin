package utils

import (
	"strings"
	"unicode"
)

// ToTitleCase returns a string with first letter uppercase and rest lowercase
func ToTitleCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
