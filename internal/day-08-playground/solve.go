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

		parent     *JunctionBox
		NrChildren int
	}

	Distance   float64
	Connection struct {
		From JunctionBox
		To   JunctionBox
	}
	Distances struct {
		BetweenBoxes map[Distance]Connection
		Lengths      []Distance
	}
)

const (
	Day string = "08"
)

func (j JunctionBox) isRoot() bool { return j.parent == nil }

func (j JunctionBox) ToS() string { return fmt.Sprintf("%d,%d,%d", j.X, j.Y, j.Z) }

func (b1 JunctionBox) DistanceTo(b2 JunctionBox) Distance {
	abs := square(b1.X-b2.X) + square(b1.Y-b2.Y) + square(b1.Z-b2.Z)

	return Distance(math.Sqrt(abs))
}

func square(a int) float64 {
	return math.Pow(float64(a), 2)
}

/**
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
*/

func FindCircuitSizes(distances Distances, firstNConnections int) []int {
	var allBoxes = make(map[string]*JunctionBox)

	for _, conn := range distances.BetweenBoxes {
		allBoxes[conn.From.ToS()] = &JunctionBox{
			X:          0,
			Y:          0,
			Z:          0,
			parent:     nil,
			NrChildren: 1,
		}
		allBoxes[conn.To.ToS()] = &JunctionBox{
			X:          0,
			Y:          0,
			Z:          0,
			parent:     nil,
			NrChildren: 1,
		}
	}

	for idx := 0; idx < firstNConnections; idx++ {
		currentLength := distances.Lengths[idx]
		currentConnection := distances.BetweenBoxes[currentLength]

		boxFrom := allBoxes[currentConnection.From.ToS()]
		boxTo := allBoxes[currentConnection.To.ToS()]

		if boxTo.isRoot() {
			boxFrom.parent = boxTo
		} else {
			boxFrom.parent = boxTo.parent
		}
		boxFrom.NrChildren += boxTo.NrChildren
	}

	allSizes := make([]int, 0)
	for _, box := range allBoxes {
		allSizes = append(allSizes, box.NrChildren)
	}

	sort.Ints(allSizes)

	return allSizes
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
				From: JunctionBox(boxes[idxFrom]),
				To:   JunctionBox(boxes[idxTo]),
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

			parent: nil,
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
	_ = boxes

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, 0)
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	boxes := ParseInput(lines)
	_ = boxes

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
