package day10factory_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-10-factory"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(Universe{}, universe)

	// Machines
	assert.Len(universe.Machines, 3)
	machine := universe.Machines[0]

	// Buttons
	assert.Equal(".##.", machine.Lights)

	// Lights
	assert.Len(machine.ButtonCombinations, 6)
	assert.Equal([]int{3}, machine.ButtonCombinations[0])
	assert.Equal([]int{0, 1}, machine.ButtonCombinations[len(machine.ButtonCombinations)-1])

	// Joltage
	assert.Equal("{3,5,4,7}", machine.Joltages)
}

func TestPressButtons(t *testing.T) {
	type Input struct {
		lights  string
		buttons []int
	}

	testCases := map[*Input]string{
		{"....", []int{3}}: "...#",
	}

	for input, expectedResult := range testCases {
		actualResult := PressButtons(input.lights, input.buttons)
		assert.Equal(t, expectedResult, actualResult)
	}
}

func testInput() []string {
	return []string{
		"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
	}
}
