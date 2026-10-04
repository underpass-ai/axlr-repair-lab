package lab

// Max returns the largest number, or 0 when there are none.
func Max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	best := 0
	for _, n := range numbers {
		if n > best {
			best = n
		}
	}
	return best
}
