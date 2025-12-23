package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/m43i/aoc-2025/internal"
)

type Calc struct {
	numbers   []int
	operation string
}

func parse(input string) []Calc {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := make([][]string, 0)

	for _, line := range lines {
		elements := strings.Fields(line)
		rows = append(rows, elements)
	}

	calc := make([]Calc, len(rows[0]))

	for _, row := range rows {
		for j, element := range row {
			if element == "*" || element == "+" {
				calc[j].operation = element
				continue
			}

			num, err := strconv.ParseInt(element, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			calc[j].numbers = append(calc[j].numbers, int(num))
		}
	}

	return calc
}

func parsePart2(input string) []Calc {
	input = strings.TrimRight(input, " \t\n\r")
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		return nil
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, maxLen)
		for j := 0; j < maxLen; j++ {
			if j < len(line) {
				grid[i][j] = rune(line[j])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	operatorRow := len(grid) - 1
	dataRows := grid[:operatorRow]

	var calcs []Calc
	var currentNumbers []int

	for col := maxLen - 1; col >= 0; col-- {
		lastRowChar := grid[operatorRow][col]

		if lastRowChar == '*' || lastRowChar == '+' {
			var numStr string
			for r := range dataRows {
				ch := dataRows[r][col]
				if ch >= '0' && ch <= '9' {
					numStr += string(ch)
				}
			}
			if numStr != "" {
				num, err := strconv.ParseInt(numStr, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				currentNumbers = append(currentNumbers, int(num))
			}

			if len(currentNumbers) > 0 {
				calcs = append(calcs, Calc{numbers: currentNumbers, operation: string(lastRowChar)})
			}
			currentNumbers = nil
			continue
		}

		if lastRowChar == ' ' {
			var numStr string
			for r := range dataRows {
				ch := dataRows[r][col]
				if ch >= '0' && ch <= '9' {
					numStr += string(ch)
				}
			}

			if numStr != "" {
				num, err := strconv.ParseInt(numStr, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				currentNumbers = append(currentNumbers, int(num))
			}
			continue
		}
	}

	return calcs
}

func calculate(calcs []Calc) int {
	result := 0
	for _, calc := range calcs {
		if calc.operation == "*" {
			result += internal.MulArray(calc.numbers)
		} else {
			result += internal.SumArray(calc.numbers)
		}
	}
	return result
}

func main() {
	rawInput := internal.GetInput(6)

	// Part 1
	input1 := parse(rawInput)
	fmt.Println("Part 1:", calculate(input1))

	// Part 2
	input2 := parsePart2(rawInput)
	fmt.Println("Part 2:", calculate(input2))
}
