package day07laboratories_test

import (
	"image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-07-laboratories"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(Universe{}, universe)
	assert.Equal(image.Pt(7, 0), universe.Start)
	assert.Equal(16, universe.MaxY)

	assert.Len(universe.Splitters, 22)
	firstSplitter := universe.Splitters[image.Pt(7, 2)]
	assert.Equal(image.Pt(7, 2), firstSplitter.Location)
	assert.Equal(0, firstSplitter.NrTimelines)
}

func TestCountSplits(t *testing.T) {
	universe := ParseInput(testInput())

	count := FollowBeamAndCountSplits(universe)

	assert.Equal(t, 21, count)
}

func TestCountTimelines(t *testing.T) {
	universe := ParseInput(testInput())

	count := FollowBeamAndCountTimelines(universe)

	assert.Equal(t, 40, count)
}

func testInput() []string {
	return []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}
}
