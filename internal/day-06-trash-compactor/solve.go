package day06trashcompactor

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "06"

type Operator uint

const (
	OperatorSum Operator = iota + 1
	OperatorMultiply
)

type (
	Column struct {
		IntRows    []int
		StringRows []string
		Operator
	}
	Grid struct {
		Columns []Column
	}
)

func (g Grid) CalculateGrandTotalString() int {
	var sum int = 0

	for _, c := range g.Columns {
		sum += c.CalculateTotalString()
	}

	return sum
}

func (g Grid) CalculateGrandTotal() int {
	var sum int = 0

	for _, c := range g.Columns {
		sum += c.CalculateTotalInt()
	}

	return sum
}

func (c Column) CalculateTotalString() int {
	var stacked []string = make([]string, len(c.StringRows[0]))

	for x := range len(c.StringRows[0]) {
		for _, raw := range c.StringRows {
			stacked[x] += string(raw[x])
		}
	}

	var nrs []int = make([]int, len(stacked))
	for idx, stack := range stacked {
		nrs[idx], _ = strconv.Atoi(strings.Trim(stack, " "))
	}

	var sum int
	switch c.Operator {
	case OperatorSum:
		sum = 0
		for _, nr := range nrs {
			sum += nr
		}
	case OperatorMultiply:
		sum = 1
		for _, nr := range nrs {
			sum *= nr
		}
	default:
		panic("No valid case found")
	}

	return sum
}

func (c Column) CalculateTotalInt() int {
	var sum int
	switch c.Operator {
	case OperatorSum:
		sum = 0
		for _, nr := range c.IntRows {
			sum += nr
		}
	case OperatorMultiply:
		sum = 1
		for _, nr := range c.IntRows {
			sum *= nr
		}
	default:
		panic("No valid case found")
	}

	return sum
}

func ParseInput(lines []string) Grid {
	var intRows [][]int = make([][]int, len(lines)-1)
	var stringRows [][]string = make([][]string, len(lines)-1)

	// Build numeric values
	for idx := range len(lines) - 1 {
		intRows[idx] = parseLineForInt(lines[idx])
	}

	// Build string values
	operatorLocations := OperatorLocations(lines[len(lines)-1])
	for idx := range len(lines) - 1 {
		stringRows[idx] = parseLineForString(lines[idx], operatorLocations)
	}

	// Build Grid
	var grid = Grid{
		Columns: make([]Column, len(intRows[0])),
	}
	// Initialize grid-columns
	for x := range grid.Columns {
		grid.Columns[x].IntRows = make([]int, len(intRows))
		grid.Columns[x].StringRows = make([]string, len(intRows))
	}

	// Populate grid columns and rows
	for y, row := range intRows {
		for x, number := range row {
			grid.Columns[x].IntRows[y] = number
			grid.Columns[x].StringRows[y] = stringRows[y][x]
		}
	}

	ex := regexp.MustCompile(`([+\*])`)
	operatorLine := lines[len(lines)-1]
	matches := ex.FindAllString(operatorLine, -1)
	fmt.Printf("matches: %v\n", matches)

	for idx, match := range matches {
		grid.Columns[idx].Operator = toOperator(match)
	}

	return grid
}

func OperatorLocations(line string) []int {
	var locations []int = make([]int, 0)

	for idx := range len(line) {
		if string(line[idx]) != " " {
			locations = append(locations, idx)
		}
	}

	return locations
}

func parseLineForString(line string, operatorLocations []int) []string {
	var values []string = make([]string, len(operatorLocations))

	for idx := range len(operatorLocations) - 1 {
		start := operatorLocations[idx]
		end := operatorLocations[idx+1] - 1

		values[idx] = line[start:end]
	}
	values[len(values)-1] = line[operatorLocations[len(operatorLocations)-1]:]

	return values
}

func parseLineForInt(line string) []int {
	ex := regexp.MustCompile(`(\d+)`)
	matches := ex.FindAllString(line, -1)

	var columns []int = make([]int, len(matches))
	for idx, match := range matches {
		columns[idx] = toInt(match)
	}

	return columns
}

func toOperator(match string) Operator {
	return map[string]Operator{
		"+": OperatorSum,
		"*": OperatorMultiply,
	}[match]
}

func toInt(match string) int {
	nr, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	grid := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, grid.CalculateGrandTotal())
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	grid := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, grid.CalculateGrandTotalString())
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
