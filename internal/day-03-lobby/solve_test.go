package day03lobby_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-03-lobby"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	banks := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal("[]day03lobby.Bank", fmt.Sprintf("%T", banks))
	assert.Len(banks, 4)
	assert.Len(banks[0].Batteries, 15)
	assert.Equal(9, banks[0].Batteries[0])
	assert.Equal(8, banks[len(banks)-1].Batteries[0])
	assert.Equal(1, banks[len(banks)-1].Batteries[len(banks[len(banks)-1].Batteries)-1])
}

func TestBankFindMaxJoltage(t *testing.T) {
	testCases := map[int]int{
		0: 98,
		1: 89,
		2: 78,
		3: 92,
	}

	banks := ParseInput(testInput())

	for idx, expectedJoltage := range testCases {
		actualJoltage := banks[idx].MaxJoltage()
		assert.Equal(t, expectedJoltage, actualJoltage)
	}
}

func TestBankFindMaxUnsafeJoltage(t *testing.T) {
	testCases := map[int]int{
		0: 987654321111,
		1: 811111111119,
		2: 434234234278,
		3: 888911112111,
	}

	banks := ParseInput(testInput())

	for idx, expectedJoltage := range testCases {
		actualJoltage := banks[idx].MaxUnsafeJoltage()
		assert.Equal(t, expectedJoltage, actualJoltage)
	}
}

func TestSumMaxJoltages(t *testing.T) {
	banks := ParseInput(testInput())

	sum := SumMaxJoltages(banks)

	assert.Equal(t, 357, sum)
}

func TestSumMaxiUnsafeJoltages(t *testing.T) {
	banks := ParseInput(testInput())

	sum := SumMaxUnsafeJoltages(banks)

	assert.Equal(t, 3121910778619, sum)
}

func testInput() []string {
	return []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
}
