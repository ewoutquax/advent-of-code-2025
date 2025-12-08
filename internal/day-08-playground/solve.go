package day08playground

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

type (
	JunctionBox struct {
		X int
		Y int
		Z int
	}

	Distance          float64
	JunctionBoxString string
	Connection        struct {
		From JunctionBoxString
		To   JunctionBoxString
	}
	Distances struct {
		BetweenBoxes map[Distance]Connection
		Lengths      []Distance
	}

	Circuit []JunctionBoxString
	Clique  []JunctionBoxString
)

const (
	Day string = "08"
)

func (j JunctionBox) ToS() string { return fmt.Sprintf("%d,%d,%d", j.X, j.Y, j.Z) }

func (b1 JunctionBox) DistanceTo(b2 JunctionBox) Distance {
	abs := square(b1.X-b2.X) + square(b1.Y-b2.Y) + square(b1.Z-b2.Z)

	return Distance(math.Sqrt(abs))
}

func square(a int) float64 {
	return math.Pow(float64(a), 2)
}

func FindMinNeededConnections(boxes []JunctionBox) int {
	distances := DistanceBetweenBoxes(boxes)

	var lowIdx int = 0
	var highIdx int = len(distances.Lengths)

	for highIdx-lowIdx > 1 {
		currentIdx := (lowIdx + highIdx) / 2

		allSizes := FindCircuits(boxes, currentIdx)

		if len(allSizes) == 1 && allSizes[0] == len(boxes) {
			highIdx = currentIdx
		} else {
			lowIdx = currentIdx
		}
	}

	fmt.Println(".")
	fmt.Printf("lowIdx: %v\n", lowIdx)
	fmt.Printf("highIdx: %v\n", highIdx)

	minLength := distances.Lengths[lowIdx]
	fmt.Printf("distances.BetweenBoxes[minLength]: %v\n", distances.BetweenBoxes[minLength])

	return 0
}

func FindCircuits(boxes []JunctionBox, firstNConnections int) []int {
	var circuits = make(map[JunctionBoxString]int)
	var newIdx = 0

	distances := DistanceBetweenBoxes(boxes)

	for idx := 0; idx < firstNConnections; idx++ {
		currentLength := distances.Lengths[idx]
		currentConnection := distances.BetweenBoxes[currentLength]

		var currentCircuitIdxFrom int = 0
		var currentCircuitIdxTo int = 0
		if idx, ok := circuits[currentConnection.From]; ok {
			currentCircuitIdxFrom = idx
		}
		if idx, ok := circuits[currentConnection.To]; ok {
			currentCircuitIdxTo = idx
		}
		if currentCircuitIdxFrom == 0 && currentCircuitIdxTo == 0 {
			newIdx++
			circuits[currentConnection.From] = newIdx
			circuits[currentConnection.To] = newIdx
		}
		if currentCircuitIdxFrom != 0 && currentCircuitIdxTo == 0 {
			circuits[currentConnection.From] = currentCircuitIdxFrom
			circuits[currentConnection.To] = currentCircuitIdxFrom
		}
		if currentCircuitIdxFrom == 0 && currentCircuitIdxTo != 0 {
			circuits[currentConnection.From] = currentCircuitIdxTo
			circuits[currentConnection.To] = currentCircuitIdxTo
		}
		if currentCircuitIdxFrom != 0 && currentCircuitIdxTo != 0 {
			lowIdx := min(currentCircuitIdxFrom, currentCircuitIdxTo)

			for name, existingIdx := range circuits {
				if existingIdx == currentCircuitIdxFrom || existingIdx == currentCircuitIdxTo {
					circuits[name] = lowIdx
				}
			}
		}
	}

	circuitSizes := make(map[int]int)
	for _, idx := range circuits {
		if count, ok := circuitSizes[idx]; ok {
			circuitSizes[idx] = count + 1
		} else {
			circuitSizes[idx] = 1
		}
	}

	allSizes := make([]int, 0)
	for _, s := range circuitSizes {
		allSizes = append(allSizes, s)
	}

	if len(allSizes) == 1 {
		fmt.Printf("circuits: %v\n", circuits)
	}

	sort.Ints(allSizes)

	return allSizes

	// return allSizes[len(allSizes)-1] *
	// 	allSizes[len(allSizes)-2] *
	// 	allSizes[len(allSizes)-3]
}

func FindCliques(boxes []JunctionBox, firstNConnections int) []Clique {
	type Neighbours map[JunctionBoxString][]JunctionBoxString
	var neighbours = make(Neighbours)

	distances := DistanceBetweenBoxes(boxes)

	for idx := range firstNConnections {
		currentDistance := distances.Lengths[idx]

		boxFrom := distances.BetweenBoxes[currentDistance].From
		boxTo := distances.BetweenBoxes[currentDistance].To

		if _, ok := neighbours[boxFrom]; !ok {
			neighbours[boxFrom] = make([]JunctionBoxString, 0)
		}
		neighbours[boxFrom] = append(neighbours[boxFrom], boxTo)

		if _, ok := neighbours[boxTo]; !ok {
			neighbours[boxTo] = make([]JunctionBoxString, 0)
		}
		neighbours[boxTo] = append(neighbours[boxTo], boxFrom)
	}
	fmt.Printf("neighbours: %v\n", neighbours)

	var initialParty = make([]JunctionBoxString, 0, len(neighbours))
	for box := range neighbours {
		initialParty = append(initialParty, box)
	}

	/**
	 * R => Current clique
	 * P => All vertices, optionally the neighbours of the current node
	 * X => Excluded vertices, to prevent duplicates
	 */
	var maxCliques []Clique = make([]Clique, 0)

	var BornKerbosch func(R Clique, P, X []JunctionBoxString)
	BornKerbosch = func(R Clique, P, X []JunctionBoxString) {
		if len(P) == 0 && len(X) == 0 {
			fmt.Printf("BornKerbosch: returning clique: %v\n", R)
			maxCliques = append(maxCliques, R)
			return
		}

		fmt.Println("")
		fmt.Println("BornKerbosch:")
		fmt.Printf("R: %v\n", R)
		fmt.Printf("P: %v\n", P)
		fmt.Printf("x: %v\n", X)

		localP := deepClone(P)
		for len(localP) > 0 {
			currentBox := localP[0]

			subR := append(R, currentBox)
			subP := Intersection(localP, neighbours[currentBox])
			subX := Intersection(X, neighbours[currentBox])

			BornKerbosch(subR, subP, subX)

			localP = localP[1:]
			X = append(X, currentBox)
		}
	}
	BornKerbosch(make(Clique, 0), initialParty, make([]JunctionBoxString, 0))

	return maxCliques
}

func deepClone(in []JunctionBoxString) []JunctionBoxString {
	var out = make([]JunctionBoxString, len(in))

	copy(out, in)

	return out
}

func Intersection(lefts, right []JunctionBoxString) []JunctionBoxString {
	var intersected []JunctionBoxString = make([]JunctionBoxString, 0, min(len(lefts), len(right)))

	for _, left := range lefts {
		if slices.Contains(right, left) {
			intersected = append(intersected, left)
		}
	}

	return intersected
}

func DistanceBetweenBoxes(boxes []JunctionBox) Distances {
	var distances = Distances{
		BetweenBoxes: make(map[Distance]Connection),
		Lengths:      make([]Distance, 0),
	}

	for idxFrom := 0; idxFrom < len(boxes)-1; idxFrom++ {
		for idxTo := idxFrom + 1; idxTo < len(boxes); idxTo++ {
			distance := boxes[idxFrom].DistanceTo(boxes[idxTo])

			if _, ok := distances.BetweenBoxes[distance]; ok {
				panic("We got the exact same distance twice!")
			}

			distances.BetweenBoxes[distance] = Connection{
				From: JunctionBoxString(boxes[idxFrom].ToS()),
				To:   JunctionBoxString(boxes[idxTo].ToS()),
			}

			distances.Lengths = append(distances.Lengths, distance)
		}
	}

	slices.Sort(distances.Lengths)

	return distances
}

func ParseInput(lines []string) []JunctionBox {
	var boxes []JunctionBox = make([]JunctionBox, len(lines))

	for idx, line := range lines {
		parts := strings.Split(line, ",")

		boxes[idx] = JunctionBox{
			X: toInt(parts[0]),
			Y: toInt(parts[1]),
			Z: toInt(parts[2]),
		}
	}

	return boxes
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
	boxes := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, FindCircuits(boxes, 1000))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	boxes := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, FindMinNeededConnections(boxes))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
