package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

func isInvalidID(min, max int) []int {
    seen := make(map[int]bool)
    var result []int

    maxDigits := len(fmt.Sprintf("%d", max))

    for patternLen := 1; patternLen <= maxDigits/2; patternLen++ {
        for reps := 2; patternLen*reps <= maxDigits; reps++ {
            patternMin := 1
            if patternLen > 1 {
                patternMin = intPow(10, patternLen-1)
            }
            patternMax := intPow(10, patternLen) - 1
            multiplier := intPow(10, patternLen)

            for pattern := patternMin; pattern <= patternMax; pattern++ {
                num := 0
                for i := 0; i < reps; i++ {
                    num = num*multiplier + pattern
                }

                if num > max {
                    break
                }
                if num >= min && !seen[num] {
                    seen[num] = true
                    result = append(result, num)
                }
            }
        }
    }

    sort.Ints(result)
    return result
}

func intPow(base, exp int) int {
    result := 1
    for range exp {
        result *= base
    }
    return result
}

func parse(input string) [][]int {
	lines := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")
	result := make([][]int, 0, len(lines))
	for _, line := range lines {
		ids := strings.Split(line, "-")
		lower, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		upper, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, []int{int(lower), int(upper)})
	}

	return result
}

func main() {
	input := parse(internal.GetInput(2))

	result := 0
	for _, r := range input {
		s := isInvalidID(r[0], r[1])
		for _, v := range s {
			result += v
		}
	}

	fmt.Println("Result:", result)
}
