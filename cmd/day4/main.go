package main

import (
	"fmt"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

var directions = [8][2]int{
	{-1, 0},  // LEFT
	{-1, -1}, // LEFT UP
	{0, -1},  // UP
	{1, -1},  // RIGHT UP
	{1, 0},   // RIGHT
	{1, 1},   // RIGHT DOWN
	{0, 1},   // DOWN
	{-1, 1},  // LEFT DOWN
}

func countAround(input [][]int, x, y int) int {
	count := 0
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if ny >= 0 && ny < len(input) && nx >= 0 && nx < len(input[ny]) {
			count += input[ny][nx]
		}
	}
	return count
}

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	result := make([][]int, len(lines)-1)
	for i, line := range lines {
		splits := strings.SplitSeq(line, "")
		for char := range splits {
			if char == "@" {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], 0)
			}
		}
	}

	return result
}

func main() {
	input := parse(internal.GetInput(4))

	result := 0
	currentInput := input

	for {
		remove := make([][2]int, 0)
		count := 0
		for y, line := range currentInput {
			for x, pos := range line {
				if pos == 1 {
					c := countAround(currentInput, x, y)
					if c < 4 {
						count++
						remove = append(remove, [2]int{x, y})
					}
				}
			}
		}

		if count > 0 {
			result += count
			for _, r := range remove {
				currentInput[r[1]][r[0]] = 0
			}
		} else {
			break
		}
	}

	fmt.Println("Count:", result)
}
