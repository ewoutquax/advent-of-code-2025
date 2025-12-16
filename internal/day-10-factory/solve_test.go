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
	assert.Equal(ButtonCombination{3}, machine.ButtonCombinations[0])
	assert.Equal(ButtonCombination{0, 1}, machine.ButtonCombinations[len(machine.ButtonCombinations)-1])

	// Joltage
	assert.Len(machine.Joltages, 4)
	assert.Equal(3, machine.Joltages[0])
	assert.Equal(5, machine.Joltages[1])
	assert.Equal(4, machine.Joltages[2])
	assert.Equal(7, machine.Joltages[3])
}

func TestPressButtonsForLights(t *testing.T) {
	type Input struct {
		lights  string
		buttons []int
	}

	testCases := map[*Input]string{
		{"....", []int{3}}:           "...#",
		{"...#", []int{3}}:           "....",
		{".#.#", []int{2}}:           ".###",
		{".#.#.", []int{1, 2, 3, 4}}: "..#.#",
	}

	for input, expectedResult := range testCases {
		actualResult := PressButtonsForLights(input.lights, input.buttons)
		assert.Equal(t, expectedResult, actualResult)
	}
}

func TestPressButtonsForJoltages(t *testing.T) {
	type Input struct {
		joltages Joltages
		buttons  []int
	}

	testCases := map[*Input]Joltages{
		{[]int{0, 0, 0, 1}, []int{3}}:          {0, 0, 0, 0},
		{[]int{3, 2, 1, 0}, []int{0, 1, 2, 3}}: []int{2, 1, 0, -1},
	}

	for input, expectedResult := range testCases {
		actualResult := PressButtonsForJoltages(input.joltages, input.buttons)
		assert.Equal(t, expectedResult, actualResult)
	}
}

func TestMachineFindNrButtonsCombinationsForLights(t *testing.T) {
	universe := ParseInput(testInput())

	testCases := map[*Machine]int{
		&universe.Machines[0]: 2,
		&universe.Machines[1]: 3,
		&universe.Machines[2]: 2,
	}

	for inputMachine, expectedNrCombinations := range testCases {
		actualNrCombinations := FindNrButtonCombinationsForLights(*inputMachine)
		assert.Equal(t, expectedNrCombinations, actualNrCombinations)
	}
}

func TestSumMinNrButtonCombinationsForLights(t *testing.T) {
	universe := ParseInput(testInput())

	sum := SumMinNrButtonCombinationsForLights(universe)

	assert.Equal(t, 7, sum)
}

func TestMachineFindNrButtonsCombinationsForJoltage(t *testing.T) {
	universe := ParseInput(testInput())

	testCases := map[*Machine]int{
		&universe.Machines[0]: 10,
		&universe.Machines[1]: 12,
		&universe.Machines[2]: 11,
	}

	for inputMachine, expectedNrCombinations := range testCases {
		actualNrCombinations := FindNrButtonCombinationsForJoltage(*inputMachine)
		assert.Equal(t, expectedNrCombinations, actualNrCombinations)
	}
}

func TestSumMinNrButtonCombinationsForJoltages(t *testing.T) {
	universe := ParseInput(testInput())

	sum := SumMinNrButtonCombinationsForJoltages(universe)

	assert.Equal(t, 33, sum)
}

func testInput() []string {
	return []string{
		"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
	}
}
