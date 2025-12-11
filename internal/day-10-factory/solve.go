package day10factory

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "10"

type (
	ButtonCombination []int
	Joltages          []int

	Machine struct {
		Lights             string
		ButtonCombinations []ButtonCombination
		CombinationRanking []int
		Joltages
	}
	Universe struct {
		Machines []Machine
	}

	PathJoltage struct {
		currentJoltages      Joltages
		NrButtonCombinations int
		TotalRanking         int
	}

	PathLight struct {
		currentLights        string
		ButtonCombinations   []ButtonCombination
		NrButtonCombinations int
	}
	PathLightHeap []PathLight
)

func (h PathLightHeap) Len() int { return len(h) }
func (h PathLightHeap) Less(i, j int) bool {
	return h[i].NrButtonCombinations < h[j].NrButtonCombinations
}
func (h PathLightHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PathLightHeap) Push(x any)   { *h = append(*h, x.(PathLight)) }

func (h *PathLightHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type PathJoltageHeap []PathJoltage

func (h PathJoltageHeap) Len() int { return len(h) }
func (h PathJoltageHeap) Less(i, j int) bool {
	return h[i].TotalRanking < h[j].TotalRanking ||
		h[i].TotalRanking == h[j].TotalRanking && h[i].NrButtonCombinations < h[j].NrButtonCombinations
}
func (h PathJoltageHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PathJoltageHeap) Push(x any)   { *h = append(*h, x.(PathJoltage)) }

func (h *PathJoltageHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (js Joltages) toHash() string {
	s := make([]string, len(js))

	for _, j := range js {
		s = append(s, fmt.Sprintf("%d", j))
	}

	return strings.Join(s, ",")
}

func (js Joltages) isDone() bool {
	for _, j := range js {
		if j != 0 {
			return false
		}
	}
	return true
}

func (js Joltages) isValid() bool {
	for _, j := range js {
		if j < 0 {
			return false
		}
	}
	return true
}

func (i ButtonCombination) eq(j ButtonCombination) bool {
	if len(i) != len(j) {
		return false
	}

	for idx := range len(i) {
		if i[idx] != j[idx] {
			return false
		}
	}

	return true
}

func PressButtonsForJoltages(joltages Joltages, buttons []int) Joltages {
	newJoltages := make([]int, len(joltages))
	copy(newJoltages, joltages)

	for _, idx := range buttons {
		newJoltages[idx] -= 1
	}

	return newJoltages
}

func PressButtonsForLights(lightString string, buttons []int) string {
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

func SumMinNrButtonCombinationsForLights(u Universe) int {
	var sum int = 0

	for _, machine := range u.Machines {
		sum += FindNrButtonCombinationsForLights(machine)
	}

	return sum
}

func SumMinNrButtonCombinationsForJoltages(u Universe) int {
	var sum int = 0

	for idx, machine := range u.Machines {
		count := FindNrButtonCombinationsForJoltage(machine)
		sum += count
		fmt.Printf("(%d/%d) count: %v\n", idx, len(u.Machines), count)
	}

	return sum
}

func FindNrButtonCombinationsForJoltage(machine Machine) int {
	visitedJoltages := make(map[string]struct{})
	var pathHeap = make(PathJoltageHeap, 0)

	heap.Init(&pathHeap)
	heap.Push(&pathHeap, PathJoltage{
		currentJoltages:      machine.Joltages,
		NrButtonCombinations: 0,
		TotalRanking:         0,
	})

	for len(pathHeap) > 0 {
		currentPath := heap.Pop(&pathHeap).(PathJoltage)
		// fmt.Printf(
		// 	"currentPath after %d steps: %v\n",
		// 	currentPath.NrButtonCombinations,
		// 	currentPath.currentJoltages,
		// )

		for idx, nextCombination := range machine.ButtonCombinations {
			newPath := PathJoltage{
				currentJoltages:      PressButtonsForJoltages(currentPath.currentJoltages, nextCombination),
				NrButtonCombinations: currentPath.NrButtonCombinations + 1,
				TotalRanking:         machine.CombinationRanking[idx],
			}

			if newPath.currentJoltages.isDone() {
				return newPath.NrButtonCombinations
			}

			hash := newPath.currentJoltages.toHash()
			if _, ok := visitedJoltages[hash]; !ok && newPath.currentJoltages.isValid() {
				visitedJoltages[hash] = struct{}{}
				heap.Push(&pathHeap, newPath)
			}
		}
	}

	panic("No solution found")
}

func FindNrButtonCombinationsForLights(machine Machine) int {
	pathHeap := make(PathLightHeap, 0)
	heap.Init(&pathHeap)
	heap.Push(&pathHeap, PathLight{
		currentLights:        strings.ReplaceAll(machine.Lights, "#", "."),
		ButtonCombinations:   make([]ButtonCombination, 0),
		NrButtonCombinations: 0,
	})

	for len(pathHeap) > 0 {
		currentPath := heap.Pop(&pathHeap).(PathLight)

		for _, nextButtonCombination := range nextValidButtonCombinations(currentPath.ButtonCombinations, machine.ButtonCombinations) {
			newLights := PressButtonsForLights(
				currentPath.currentLights,
				nextButtonCombination,
			)

			if newLights == machine.Lights {
				return currentPath.NrButtonCombinations + 1
			}

			newPath := PathLight{
				currentLights:        newLights,
				ButtonCombinations:   append(currentPath.ButtonCombinations, nextButtonCombination),
				NrButtonCombinations: currentPath.NrButtonCombinations + 1,
			}

			heap.Push(&pathHeap, newPath)
		}
	}

	panic("No solution found")
}

func nextValidButtonCombinations(usedCombinations []ButtonCombination, allCombinations []ButtonCombination) []ButtonCombination {
	var nextCombinations []ButtonCombination = make([]ButtonCombination, 0)

	for _, possibleCombination := range allCombinations {
		if !slices.ContainsFunc(usedCombinations, func(usedCombination ButtonCombination) bool {
			return usedCombination.eq(possibleCombination)
		}) {
			nextCombinations = append(nextCombinations, possibleCombination)
		}
	}

	return nextCombinations
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

	buttons := make([]ButtonCombination, len(parts)-2)
	for idx := 1; idx < len(parts)-1; idx++ {
		raw := parts[idx][1 : len(parts[idx])-1]

		localButtons := make([]int, 0)
		for s := range strings.SplitSeq(raw, ",") {
			localButtons = append(localButtons, toInt(s))
		}

		buttons[idx-1] = localButtons
	}

	groupedButtons := make(map[int]int)
	for _, localButtons := range buttons {
		for _, btn := range localButtons {
			if count, ok := groupedButtons[btn]; ok {
				groupedButtons[btn] = count + 1
			} else {
				groupedButtons[btn] = 1
			}
		}
	}

	combinationRanking := make([]int, len(buttons))
	for idx, localButtons := range buttons {
		var ranking int = 999999
		for _, nr := range localButtons {
			ranking = min(ranking, groupedButtons[nr])
		}
		combinationRanking[idx] = ranking
	}

	fmt.Printf("groupedButtons: %v\n", groupedButtons)
	fmt.Printf("combinationRanking: %v\n", combinationRanking)

	partJoltages := parts[len(parts)-1]
	partJoltages = partJoltages[1 : len(partJoltages)-1]

	joltages := make([]int, strings.Count(partJoltages, ",")+1)
	sum := 0
	for idx, s := range strings.Split(partJoltages, ",") {

		joltages[idx] = toInt(s)
		sum += joltages[idx]
	}

	return Machine{
		Lights:             lights,
		ButtonCombinations: buttons,
		Joltages:           joltages,
		CombinationRanking: combinationRanking,
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
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumMinNrButtonCombinationsForLights(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumMinNrButtonCombinationsForJoltages(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
