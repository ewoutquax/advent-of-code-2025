package day05cafetaria_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-05-cafetaria"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	freshRanges := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal("[]day05cafetaria.FreshRange", fmt.Sprintf("%T", freshRanges))
	assert.Len(freshRanges, 4)
	assert.Equal(IngredientId(3), freshRanges[0].From)
	assert.Equal(IngredientId(5), freshRanges[0].Until)
	assert.Equal(IngredientId(18), freshRanges[len(freshRanges)-1].Until)
}

func TestIngredientIsFresh(t *testing.T) {
	freshRanges := ParseInput(testInput())

	testCases := map[IngredientId]bool{
		1:  false,
		5:  true,
		8:  false,
		11: true,
		17: true,
		32: false,
	}

	for inputIngredientId, expectedResult := range testCases {
		actualResult := inputIngredientId.IsFresh(freshRanges)
		assert.Equal(t, expectedResult, actualResult, inputIngredientId)
	}
}

func TestCountFreshIngredients(t *testing.T) {
	freshRanges := ParseInput(testInput())

	count := CountFreshRanges(testInput()[1], freshRanges)

	assert.Equal(t, count, 3)
}

func TestCountFreshSize(t *testing.T) {
	freshRanges := ParseInput(testInput())

	count := CountFreshSize(freshRanges)

	assert.Equal(t, 14, count)
}

func testInput() [][]string {
	return [][]string{
		{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
		}, {
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		},
	}
}
