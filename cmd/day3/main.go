package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

func findLargestN(input []int, n int) int64 {
	result := make([]int, 0, n)
	start := 0

	for i := range n {
		end := len(input) - (n - i - 1)

		bestPos := start
		for j := start; j < end; j++ {
			if input[j] > input[bestPos] {
				bestPos = j
			}
			if input[bestPos] == 9 {
				break
			}
		}

		result = append(result, input[bestPos])
		start = bestPos + 1
	}

	var numStr strings.Builder
	for _, d := range result {
		fmt.Fprintf(&numStr, "%d", d)
	}
	num, err := strconv.ParseInt(numStr.String(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return num
}

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	result := make([][]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		chars := strings.SplitSeq(line, "")
		for char := range chars {
			num, err := strconv.ParseInt(char, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			result[i] = append(result[i], int(num))
		}
	}

	return result
}

func main() {
	input := parse(internal.GetInput(3))
	var total int64
	for _, line := range input {
		total += findLargestN(line, 12)
	}

	fmt.Println(total)
}
