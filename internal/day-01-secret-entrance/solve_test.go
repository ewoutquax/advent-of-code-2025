package day01secretentrance_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-01-secret-entrance"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	turns := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal("[]day01secretentrance.Turn", fmt.Sprintf("%T", turns))
	assert.Len(turns, 10)
	assert.Equal(DirectionLower, turns[0].Direction)
	assert.Equal(68, turns[0].Amount)
	assert.Equal(DirectionHigher, turns[2].Direction)
}

func TestCountDailAtZero(t *testing.T) {
	turns := ParseInput(testInput())

	count := CountDailAtZero(turns)
	assert.Equal(t, 3, count)
}

func TestCountDailPassesZero(t *testing.T) {
	turns := ParseInput(testInput())

	count := CountDailPassesZero(turns)
	assert.Equal(t, 6, count)
}

func testInput() []string {
	return []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
}
