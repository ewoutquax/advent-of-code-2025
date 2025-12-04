package day04printingdepartment

import (
	"fmt"
	"image"
	"slices"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "04"

type (
	PaperRoll struct {
		Location   image.Point
		Neighbours []*PaperRoll
		isRemoved  bool
	}
	PaperRolls map[string]*PaperRoll
)

func (p *PaperRoll) cleanRemovedNeighbours() {
	if !p.isRemoved {
		p.Neighbours =
			slices.DeleteFunc(
				p.Neighbours,
				func(n *PaperRoll) bool { return n.isRemoved },
			)
	}
}

func CountRemoveableRollsIterative(paperRolls PaperRolls) int {
	var count int = 0

	var subCount int = 1
	for subCount > 0 {
		for _, r := range paperRolls {
			r.cleanRemovedNeighbours()
		}

		subCount = CountMoveableRolls(paperRolls)
		count += subCount
	}

	return count
}

func CountMoveableRolls(paperRolls PaperRolls) int {
	var count int = 0

	for _, roll := range paperRolls {
		if !roll.isRemoved && len(roll.Neighbours) < 4 {
			count++
			roll.isRemoved = true
		}
	}

	return count
}

func ParseInput(lines []string) PaperRolls {
	rolls := make([]PaperRoll, 0)

	// Make all rolls
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "@" {
				rolls = append(rolls, PaperRoll{
					Location:   image.Pt(x, y),
					Neighbours: make([]*PaperRoll, 0, 8),
					isRemoved:  false,
				})
			}
		}
	}

	// Put all rolls in indexed map
	var paperRolls = make(PaperRolls, len(rolls))
	for _, roll := range rolls {
		paperRolls[roll.Location.String()] = &roll
	}

	// Link neighbours
	for _, paperRoll := range paperRolls {
		for _, vector := range vectors() {
			otherLoc := paperRoll.Location.Add(vector)
			if neighbour, ok := paperRolls[otherLoc.String()]; ok {
				paperRoll.Neighbours = append(paperRoll.Neighbours, neighbour)
			}
		}
	}

	return paperRolls
}

func vectors() []image.Point {
	return []image.Point{
		image.Pt(-1, -1),
		image.Pt(-1, 0),
		image.Pt(-1, 1),
		image.Pt(0, -1),
		image.Pt(0, 1),
		image.Pt(1, -1),
		image.Pt(1, 0),
		image.Pt(1, 1),
	}
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	paperRolls := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountMoveableRolls(paperRolls))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	paperRolls := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountRemoveableRollsIterative(paperRolls))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
