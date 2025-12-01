package day01secretentrance

import (
	"fmt"
	"strconv"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

type (
	Direction uint

	Turn struct {
		Direction
		Amount int
	}
)

const (
	DirectionLower Direction = iota + 1
	DirectionHigher

	Day       string = "01"
	MAX_DIAL  int    = 100
	INIT_DIAL int    = 50
)

func CountDailPassesZero(turns []Turn) int {
	var count int = 0

	dial := INIT_DIAL
	for _, turn := range turns {
		count += turn.Amount / MAX_DIAL
		amt := turn.Amount % MAX_DIAL

		switch turn.Direction {
		case DirectionLower:
			if amt >= dial && dial != 0 {
				count++
			}
			dial = (dial - amt + MAX_DIAL) % MAX_DIAL
		case DirectionHigher:
			if amt >= (MAX_DIAL - dial) {
				count++
			}
			dial = (dial + amt) % MAX_DIAL
		default:
			panic("No valid case found")
		}
	}

	return count
}

func CountDailAtZero(turns []Turn) int {
	var count int = 0

	dial := INIT_DIAL
	for _, turn := range turns {
		switch turn.Direction {
		case DirectionLower:
			dial = (dial - turn.Amount) % MAX_DIAL
		case DirectionHigher:
			dial = (dial + turn.Amount) % MAX_DIAL
		default:
			panic("No valid case found")
		}

		if dial == 0 {
			count++
		}
	}

	return count
}

func ParseInput(lines []string) []Turn {
	var turns []Turn = make([]Turn, len(lines))

	for idx, line := range lines {
		turns[idx] = Turn{
			Direction: toDirection(string(line[0])),
			Amount:    convAToI(line[1:]),
		}
	}

	return turns
}

func convAToI(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func toDirection(s string) Direction {
	return map[string]Direction{
		"L": DirectionLower,
		"R": DirectionHigher,
	}[s]
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	turns := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountDailAtZero(turns))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	turns := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountDailPassesZero(turns))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
