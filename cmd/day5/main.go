package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

func isNumFresh(ranges [][2]int, num int) bool {
	if len(ranges) == 0 {
		return true
	}

	for _, rang := range ranges {
		if num >= rang[0] && num <= rang[1] {
			return true
		}
	}

	return false
}

func countFreshInRanges(ranges [][2]int) int {
	if len(ranges) == 0 {
		return 0
	}

	sorted := make([][2]int, len(ranges))
	copy(sorted, ranges)

	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j][0] < sorted[i][0] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	count := 0
	current := sorted[0]

	for i := 1; i < len(sorted); i++ {
		next := sorted[i]
		if next[0] <= current[1]+1 {
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			count += current[1] - current[0] + 1
			current = next
		}
	}
	count += current[1] - current[0] + 1

	return count
}

func parse(input string) ([][2]int, []int) {
	lines := strings.Split(input, "\n")
	secondPart := false

	ranges := make([][2]int, 0)
	numbers := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			secondPart = true
			continue
		}

		if !secondPart {
			split := strings.Split(line, "-")
			first, err := strconv.ParseInt(split[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.ParseInt(split[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			ranges = append(ranges, [2]int{int(first), int(second)})
		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, int(num))
		}
	}

	return ranges, numbers
}

func main() {
	rangs, nums := parse(internal.GetInput(5))

	count := 0

	for _, num := range nums {
		if isNumFresh(rangs, num) {
			count++
		}
	}
	count = countFreshInRanges(rangs)

	fmt.Println(count)
}
