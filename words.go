// Package lab is a disposable repository the AXLR console repairs in its own
// tests. Nothing here is a product.
package lab

import "strings"

// WordCount counts the words in text separated by spaces.
func WordCount(text string) int {
	return len(strings.Split(text, " "))
}
