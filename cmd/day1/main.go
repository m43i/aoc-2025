package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

func turndial(dial int, steps int) (int, int) {
    if steps == 0 {
        return dial, 0
    }

    var finalPos int
    var count int

    if steps > 0 {
        finalPos = (dial + steps) % 100
        count = (dial + steps) / 100
    } else {
        absSteps := -steps
        finalPos = ((dial-absSteps)%100 + 100) % 100

        if dial == 0 {
            count = absSteps / 100
        } else if absSteps < dial {
            count = 0
        } else {
            count = 1 + (absSteps-dial)/100
        }
    }

    return finalPos, count
}

func parse(input string) []int {
	content := strings.ReplaceAll(input, "R", "+")
	content = strings.ReplaceAll(content, "L", "-")
	lines := strings.Split(content, "\n")

	result := make([]int, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		result[i] = int(value)
	}

	return result
}

func main() {
	turns := parse(internal.GetInput(1))
	start := 50
	count := 0
	for _, turn := range turns {
		s, c := turndial(start, turn)
		count += c
		start = s
	}	

	fmt.Printf("Final position: %d\n", start)
	fmt.Printf("Count of 0s: %d\n", count)
}
