package day11reactor_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-11-reactor"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	universe := ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(Universe{}, universe)
	assert.Len(universe.Wires, 10)

	assert.Len(universe.Wires["aaa"].To, 2)

	assert.Len(universe.Wires["iii"].To, 1)
	assert.Equal(Name("out"), universe.Wires["iii"].To[0].From)

	assert.Equal(Name("you"), universe.StartWire.From)
	assert.Equal(-1, universe.StartWire.NrStepsToOut)
}

func TestCountPathsToEnd(t *testing.T) {
	universe := ParseInput(testInput())

	count := CountPathsToOut(universe)

	assert.Equal(t, 5, count)
}

func TestCountValidPathsFromSrvToOut(t *testing.T) {
	universe := ParseInput(testInputPart2())

	count := CountValidPathsFromSrvToOut(universe)

	assert.Equal(t, 2, count)
}

func testInput() []string {
	return []string{
		"aaa: you hhh",
		"you: bbb ccc",
		"bbb: ddd eee",
		"ccc: ddd eee fff",
		"ddd: ggg",
		"eee: out",
		"fff: out",
		"ggg: out",
		"hhh: ccc fff iii",
		"iii: out",
	}
}

func testInputPart2() []string {
	return []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}
}
