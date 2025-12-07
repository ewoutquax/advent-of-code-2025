package day07laboratories

import (
	"fmt"
	"image"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "07"

type (
	Splitter struct {
		Location    image.Point
		NrTimelines int
	}
	Universe struct {
		Start     image.Point
		Splitters map[image.Point]*Splitter
		MaxY      int
	}
)

func (s *Splitter) GetNrTimelines(u Universe) int {
	if s.NrTimelines != 0 {
		return s.NrTimelines
	}

	timelinesBeamL := FollowSingleBeamAndCountTimelines(image.Pt(s.Location.X-1, s.Location.Y), u)
	timelinesBeamR := FollowSingleBeamAndCountTimelines(image.Pt(s.Location.X+1, s.Location.Y), u)

	s.NrTimelines = timelinesBeamL + timelinesBeamR

	return s.NrTimelines
}

func FollowBeamAndCountTimelines(u Universe) int {
	return FollowSingleBeamAndCountTimelines(u.Start, u)
}

func FollowSingleBeamAndCountTimelines(beam image.Point, u Universe) int {
	for y := beam.Y; y < u.MaxY; y++ {
		if splitter, ok := u.Splitters[image.Pt(beam.X, y)]; ok {
			return splitter.GetNrTimelines(u)
		}
	}

	return 1
}

func FollowBeamAndCountSplits(u Universe) int {
	var count int = 0
	var beams = make(map[int]struct{})

	beams[u.Start.X] = struct{}{}
	for y := range u.MaxY {
		newBeams := make(map[int]struct{})

		for locX := range beams {
			if _, ok := u.Splitters[image.Pt(locX, y)]; ok {
				newBeams[locX-1] = struct{}{}
				newBeams[locX+1] = struct{}{}

				count++
			} else {
				newBeams[locX] = struct{}{}
			}
		}

		beams = newBeams
	}

	return count
}

func ParseInput(lines []string) Universe {
	var universe = Universe{
		Start:     image.Point{},
		Splitters: make(map[image.Point]*Splitter),
		MaxY:      len(lines),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			switch char {
			case "S":
				universe.Start = image.Pt(x, y)
			case "^":
				universe.Splitters[image.Pt(x, y)] = &Splitter{
					Location:    image.Pt(x, y),
					NrTimelines: 0,
				}

			case ".":
				continue
			default:
				panic("No valid case found")
			}
		}
	}

	return universe
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, FollowBeamAndCountSplits(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, FollowBeamAndCountTimelines(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
