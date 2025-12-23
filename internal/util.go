package internal

func SumArray(input []int) int {
	result := 0
	for _, num := range input {
		result += num
	}

	return result
}

func MulArray(input []int) int {
	result := 1
	for _, num := range input {
		result *= num
	}

	return result
}
