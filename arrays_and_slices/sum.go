package arraysandslices

func Sum(numbers [4]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
