package day09movietheater_test

import (
	"image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-09-movie-theater"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	tiles := ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(tiles, 8)
	assert.Equal(Tile(image.Pt(7, 1)), tiles[0])
	assert.Equal(Tile(image.Pt(7, 3)), tiles[len(tiles)-1])
}

func TestFindBiggestRectangle(t *testing.T) {
	tiles := ParseInput(testInput())

	area := FindBiggestRectangle(tiles)

	assert.Equal(t, 50, area)
}

func testInput() []string {
	return []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}
}
