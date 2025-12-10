package day10factory

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "10"

type (
	Machine struct {
		Lights             string
		ButtonCombinations [][]int
		Joltages           string
	}
	Universe struct {
		Machines []Machine
	}
	Path struct {
		currentLights        string
		ButtonCombinations   [][]int
		NrButtonCombinations int
	}
)

type PathHeap []Path

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].NrButtonCombinations < h[j].NrButtonCombinations }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x any)        { *h = append(*h, x.(Path)) }

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func PressButtons(lightString string, buttons []int) string {
	var lights = strings.Split(lightString, "")

	for _, idx := range buttons {
		if lights[idx] == "." {
			lights[idx] = "#"
		} else {
			lights[idx] = "."
		}
	}

	return strings.Join(lights, "")
}

func FindNrButtonCombinations(machine Machine) int {
	paths := make([]Path, 0)
	paths = append(paths, Path{})

	pathHeap = make(pathHeap, 0)

	return 0
}

func ParseInput(lines []string) Universe {
	var machines = make([]Machine, len(lines))

	for idx, line := range lines {
		machines[idx] = parseMachine(line)
	}

	return Universe{
		Machines: machines,
	}
}

func parseMachine(line string) Machine {
	parts := strings.Split(line, " ")

	lights := parts[0][1 : len(parts[0])-1]

	buttons := make([][]int, len(parts)-2)
	for idx := 1; idx < len(parts)-1; idx++ {
		raw := parts[idx][1 : len(parts[idx])-1]

		localButtons := make([]int, 0)
		for _, s := range strings.Split(raw, ",") {
			localButtons = append(localButtons, toInt(s))
		}

		buttons[idx-1] = localButtons
	}

	return Machine{
		Lights:             lights,
		ButtonCombinations: buttons,
		Joltages:           parts[len(parts)-1],
	}
}

func toInt(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	_ = lines

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, 0)
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	_ = lines

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}
