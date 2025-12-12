package day12christmastreefarm_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-12-christmas-tree-farm"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Shapes, 6)
	assert.Len(universe.Areas, 3)

	firstShape := universe.Shapes[0]
	assert.Equal(7, firstShape.NrBlocks)

	firstArea := universe.Areas[0]
	assert.Equal(4, firstArea.Dimensions.Dx())
	assert.Equal(4, firstArea.Dimensions.Dy())
	assert.Len(firstArea.Packages, 6)
	assert.Equal(ShapeQuantity(0), firstArea.Packages[0])
	assert.Equal(ShapeQuantity(2), firstArea.Packages[4])
}

func TestAreaCanPossibleFitTheirPackages(t *testing.T) {
	universe := ParseInput(testInput())

	testCases := map[*Area]bool{
		&universe.Areas[0]: true,
		&universe.Areas[1]: true,
		&universe.Areas[2]: true,
	}

	for inputArea, expectedResult := range testCases {
		actualResult := inputArea.CanPossibleFitTheirPackages(universe.Shapes)

		assert.Equal(t, expectedResult, actualResult)
	}
}

func TestCountAreaCanPossibleFitTheirPackages(t *testing.T) {
	universe := ParseInput(testInput())

	count := CountAreaCanPossibleFitTheirPackages(universe)

	assert.Equal(t, count, 3)
}

func testInput() [][]string {
	return [][]string{
		{
			"0:",
			"###",
			"##.",
			"##.",
		}, {
			"1:",
			"###",
			"##.",
			".##",
		}, {
			"2:",
			".##",
			"###",
			"##.",
		}, {
			"3:",
			"##.",
			"###",
			"##.",
		}, {
			"4:",
			"###",
			"#..",
			"###",
		}, {
			"5:",
			"###",
			".#.",
			"###",
		}, {
			"4x4: 0 0 0 0 2 0",
			"12x5: 1 0 1 0 2 2",
			"12x5: 1 0 1 0 3 2",
		},
	}
}
