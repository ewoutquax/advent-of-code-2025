package day12christmastreefarm

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "12"

type (
	ShapeQuantity int
	Shape         struct {
		NrBlocks int
	}
	Area struct {
		Dimensions image.Rectangle
		Packages   []ShapeQuantity
	}
	Universe struct {
		Shapes []Shape
		Areas  []Area
	}
)

func (a Area) CanPossibleFitTheirPackages(shapes []Shape) bool {
	var minSpace int = 0
	for idx, quantity := range a.Packages {
		minSpace += shapes[idx].NrBlocks * int(quantity)
	}

	return a.Dimensions.Dx()*a.Dimensions.Dy() >= minSpace
}

func CountAreaCanPossibleFitTheirPackages(universe Universe) int {
	var count int = 0
	for _, area := range universe.Areas {
		if area.CanPossibleFitTheirPackages(universe.Shapes) {
			count++
		}
	}

	return count
}

func ParseInput(blocks [][]string) Universe {
	var shapes = make([]Shape, len(blocks)-1)
	for idx := 0; idx < len(blocks)-1; idx++ {
		shapes[idx] = parseShape(blocks[idx])
	}

	var areas = make([]Area, len(blocks[len(blocks)-1]))
	for idx, line := range blocks[len(blocks)-1] {
		areas[idx] = parseArea(line)
	}

	return Universe{
		Shapes: shapes,
		Areas:  areas,
	}
}

func parseArea(line string) Area {
	parts := strings.Split(line, ": ")
	dimensionParts := strings.Split(parts[0], "x")

	packageParts := strings.Split(parts[1], " ")
	var packages = make([]ShapeQuantity, len(packageParts))
	for idx, quantity := range packageParts {
		packages[idx] = ShapeQuantity(toInt(quantity))
	}

	return Area{
		Dimensions: image.Rect(0, 0, toInt(dimensionParts[0]), toInt(dimensionParts[1])),
		Packages:   packages,
	}
}

func toInt(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return nr
}

func parseShape(lines []string) Shape {
	var count int = 0
	for idx := 1; idx < len(lines); idx++ {
		for char := range strings.SplitSeq(lines[idx], "") {
			if char == "#" {
				count++
			}
		}
	}

	return Shape{
		NrBlocks: count,
	}
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountAreaCanPossibleFitTheirPackages(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	_ = lines

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}
