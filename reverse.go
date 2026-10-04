package lab

// Reverse returns text with its runes in reverse order.
func Reverse(text string) string {
	bytes := []byte(text)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
