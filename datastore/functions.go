package datastore

import (
	"strings"
)

// ToLowerContains will lower case the inputs and do a case insensitive match
func ToLowerContains(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}
