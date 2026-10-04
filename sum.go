package lab

// Sum adds the numbers.
func Sum(numbers ...int) int {
	total := 0
	for i := 1; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}
