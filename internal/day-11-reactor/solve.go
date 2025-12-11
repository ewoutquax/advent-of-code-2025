package day11reactor

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "11"

type (
	Name string
	Wire struct {
		From         Name
		To           []*Wire
		NrStepsToFft int
		NrStepsToDac int
		NrStepsToOut int
	}
	Universe struct {
		StartWire *Wire
		Wires     map[Name]*Wire
	}
)

func CountValidPathsFromSrvToOut(u Universe) int {
	count := countPathsFromWireToFft(u.Wires["svr"]) *
		countPathsFromWireToDac(u.Wires["fft"]) *
		countPathsFromWireToOut(u.Wires["dac"])

	for _, w := range u.Wires {
		w.NrStepsToFft = -1
		w.NrStepsToDac = -1
		w.NrStepsToOut = -1
	}

	return count + countPathsFromWireToDac(u.Wires["svr"])*
		countPathsFromWireToFft(u.Wires["dac"])*
		countPathsFromWireToOut(u.Wires["fft"])
}

func CountPathsToOut(u Universe) int {
	return countPathsFromWireToOut(u.StartWire)
}

func countPathsFromWireToDac(wire *Wire) int {
	if wire.NrStepsToDac != -1 {
		return wire.NrStepsToDac
	}
	if wire.From == "dac" {
		return 1
	}
	if len(wire.To) == 0 {
		return 0
	}

	var count int = 0
	for _, w := range wire.To {
		count += countPathsFromWireToDac(w)
	}
	wire.NrStepsToDac = count

	return count
}

func countPathsFromWireToFft(wire *Wire) int {
	if wire.NrStepsToFft != -1 {
		return wire.NrStepsToFft
	}
	if wire.From == "fft" {
		return 1
	}
	if len(wire.To) == 0 {
		return 0
	}

	var count int = 0
	for _, w := range wire.To {
		count += countPathsFromWireToFft(w)
	}
	wire.NrStepsToFft = count

	return count
}

func countPathsFromWireToOut(wire *Wire) int {
	if wire.NrStepsToOut != -1 {
		return wire.NrStepsToOut
	}
	if wire.From == "out" {
		return 1
	}

	var count int = 0
	for _, w := range wire.To {
		count += countPathsFromWireToOut(w)
	}
	wire.NrStepsToOut = count

	return count
}

func ParseInput(lines []string) Universe {
	var universe Universe = Universe{
		Wires:     make(map[Name]*Wire),
		StartWire: &Wire{},
	}

	// Find all wires
	var allWires = make(map[Name]*Wire)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		allWires[Name(parts[0])] = &Wire{
			From:         Name(parts[0]),
			To:           make([]*Wire, 0),
			NrStepsToOut: -1,
			NrStepsToFft: -1,
			NrStepsToDac: -1,
		}

		for _, toPart := range strings.Split(parts[1], " ") {
			allWires[Name(toPart)] = &Wire{
				From:         Name(toPart),
				To:           make([]*Wire, 0),
				NrStepsToOut: -1,
				NrStepsToFft: -1,
				NrStepsToDac: -1,
			}
		}
	}

	// Connect all wires
	for _, line := range lines {
		parts := strings.Split(line, ": ")

		fromWire := allWires[Name(parts[0])]
		for _, toPart := range strings.Split(parts[1], " ") {
			fromWire.To = append(fromWire.To, allWires[Name(toPart)])
		}
		universe.Wires[fromWire.From] = fromWire

		if fromWire.From == "you" {
			universe.StartWire = fromWire
		}
	}

	return universe
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountPathsToOut(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountValidPathsFromSrvToOut(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
