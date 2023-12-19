package utils

import "strings"

// Transforms a sentence to all caps with an exclamation point
func MakeExcited(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}