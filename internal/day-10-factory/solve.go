package day10factory

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
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
		count, _ := FindNrButtonCombinationsForJoltage(machine)
		sum += count
		fmt.Printf("(%d/%d) count: %v\n", idx, len(u.Machines), count)
	}

	return sum
}

func FindNrButtonCombinationsForJoltage(machine Machine) (int, error) {
	fmt.Println(".")
	fmt.Println("FindNrButtonCombinationsForJoltage: start")
	fmt.Printf("machine.Joltages: %v\n", machine.Joltages)

	if machine.Joltages.isDone() {
		return 0, nil
	}

	if !machine.Joltages.isValid() {
		return 0, errors.New("Invalid joltages received")
	}

	if len(machine.ButtonCombinations) == 0 {
		return 0, errors.New("No combinations remaining")
	}

	var lowestJoltage int = math.MaxInt
	var idxLowestJoltage int = 0

	// Find the lowest joltage
	for idx, joltage := range machine.Joltages {
		if lowestJoltage > joltage && joltage != 0 {
			lowestJoltage = joltage
			idxLowestJoltage = idx
		}
	}
	fmt.Printf("lowestJoltage: %v at button %d\n", lowestJoltage, idxLowestJoltage)

	// Find the combinations using this button
	var applicableCombinations = make([]ButtonCombination, 0)
	var remainingCombinations = make([]ButtonCombination, 0)
	for _, combination := range machine.ButtonCombinations {
		if slices.Contains(combination, idxLowestJoltage) {
			applicableCombinations = append(applicableCombinations, combination)
		} else {
			remainingCombinations = append(remainingCombinations, combination)
		}
	}
	fmt.Printf("applicableCombinations (combinations using this button): %v\n", applicableCombinations)

	// This variable holds the amount each applicable combination is used, and should always sum to the lowestJoltage
	var currentDistribution = make([]int, len(applicableCombinations))
	currentDistribution[0] = lowestJoltage

	// Internal function to find all applicable combinations
	var increaseAtIndex func(int, int) int
	increaseAtIndex = func(idx int, toSpend int) int {
		if currentDistribution[idx] > lowestJoltage || currentDistribution[idx] > toSpend {
			currentDistribution[idx] = 0
			localToSpend := toSpend + currentDistribution[idx]
			remainingToSpend := increaseAtIndex(idx, localToSpend)

			return remainingToSpend
		}
		currentDistribution[idx]++

		return toSpend - 1
	}

	// Iterate through all possible distributions of combinations that sum to the lowest joltage
	var minNrSteps int = math.MaxInt
	var isSolutionFound bool = false
	for currentDistribution[len(currentDistribution)-1] < lowestJoltage {
		fmt.Printf("currentDistribution (how many times are we using each of the applicableCombinations): %v\n", currentDistribution)

		// Apply the combinations in the currentCombination, with their current number of occurences
		newMachine := Machine{
			Lights:             "",
			ButtonCombinations: remainingCombinations,
			Joltages:           deepCopy(machine.Joltages),
		}
		for idxCombination, count := range currentDistribution {
			fmt.Printf(
				"Apply combination %d (applicableCombinations[idxCombination] %v) %d times\n",
				idxCombination,
				applicableCombinations[idxCombination],
				count,
			)

			for _, idxButton := range applicableCombinations[idxCombination] {
				newMachine.Joltages[idxButton] -= count
			}
		}

		// Recursively call this function until no solution is possible, or a solution is found
		nrSubSteps, err := FindNrButtonCombinationsForJoltage(newMachine)
		if err != nil {
			minNrSteps = min(minNrSteps, nrSubSteps)
			isSolutionFound = true
		}

		// Determine the next distribution of applicable combinations
		remaining := lowestJoltage
		if len(applicableCombinations) > 0 {
			// There are multiple combinations; let's increase each of them in turn
			toSpend := lowestJoltage
			for idx := 1; idx < len(currentDistribution); idx++ {
				toSpend -= currentDistribution[idx]
			}
			remaining = increaseAtIndex(1, toSpend)
		}
		currentDistribution[0] = remaining
	}

	// If we found a solution, return the number of combinations we used
	// Which is equal to the lowestJoltage plus the lowest number of substeps
	if isSolutionFound {
		return lowestJoltage + minNrSteps, nil
	}

	return 0, errors.New("No solution possible")
}

func deepCopy(joltages Joltages) Joltages {
	var newJoltages = make(Joltages, len(joltages))

	copy(newJoltages, joltages)

	return newJoltages
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
