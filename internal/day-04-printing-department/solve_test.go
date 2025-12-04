package day04printingdepartment_test

import (
	"image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-04-printing-department"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	paperRolls := ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(PaperRolls{}, paperRolls)
	assert.Len(paperRolls, 71)
	assert.Contains(paperRolls, image.Pt(0, 1).String())
	assert.Contains(paperRolls, image.Pt(8, 9).String())
}

func TestCountMoveableRolls(t *testing.T) {
	paperRolls := ParseInput(testInput())

	count := CountMoveableRolls(paperRolls)
	assert.Equal(t, 13, count)
}

func TestCountMoveableRollsIterative(t *testing.T) {
	paperRolls := ParseInput(testInput())

	count := CountRemoveableRollsIterative(paperRolls)
	assert.Equal(t, 43, count)
}

func testInput() []string {
	return []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
}
