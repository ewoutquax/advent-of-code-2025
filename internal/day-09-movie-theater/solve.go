package day09movietheater

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "09"

type (
	Tile image.Point
)

func FindBiggestRectangle(tiles []Tile) int {
	var maxArea int = 0

	for idx1 := 0; idx1 < len(tiles)-1; idx1++ {
		for idx2 := idx1 + 1; idx2 < len(tiles); idx2++ {
			rectangle := image.Rect(tiles[idx1].X, tiles[idx1].Y, tiles[idx2].X, tiles[idx2].Y)
			area := (rectangle.Size().X + 1) * (rectangle.Size().Y + 1)

			// fmt.Printf( "tiles[idx1] / tiles[idx2] / rect / area: %s / %s / %s / %d\n", image.Point(tiles[idx1]).String(), image.Point(tiles[idx2]).String(), rectangle.Size(), area,)

			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func ParseInput(lines []string) []Tile {
	var tiles = make([]Tile, len(lines))

	for idx, line := range lines {
		parts := strings.Split(line, ",")
		tiles[idx] = Tile(image.Pt(toInt(parts[0]), toInt(parts[1])))
	}

	return tiles
}

func toInt(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	tiles := ParseInput(lines)

	/**
	Answers:
	--------
	9699684269 : incorrect
	*/

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, FindBiggestRectangle(tiles))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	tiles := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, FindBiggestRectangle(tiles))
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)

}
