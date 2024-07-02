package iteration

import "strings"

// Repeat repeats given character given amount of times
func Repeat(character string, repeatAmount int) string {
	return strings.Repeat(character, repeatAmount)
}
