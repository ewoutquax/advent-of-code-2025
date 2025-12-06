package day05cafetaria

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "05"

type (
	IngredientId int
	FreshRange   struct {
		From  IngredientId
		Until IngredientId
	}
)

func (i FreshRange) deepCopy() FreshRange {
	return FreshRange{
		From:  i.From,
		Until: i.Until,
	}
}

func (i IngredientId) IsFresh(freshRanges []FreshRange) bool {
	for _, freshRange := range freshRanges {
		if freshRange.From <= i && freshRange.Until >= i {
			return true
		}
	}
	return false
}

func CountFreshSize(freshRanges []FreshRange) int {
	slices.SortFunc(freshRanges, func(i, j FreshRange) int {
		if i.From < j.From ||
			i.From == j.From &&
				i.Until < j.Until {
			return -1
		}
		return 1
	})

	var count int = 0

	for idx, freshRange := range freshRanges {
		currentFreshRange := freshRange.deepCopy()
		for offset := range idx {
			if currentFreshRange.From >= freshRanges[offset].From && currentFreshRange.From <= freshRanges[offset].Until {
				currentFreshRange.From = freshRanges[offset].Until + 1
			}
			if currentFreshRange.Until >= freshRanges[offset].From && currentFreshRange.Until <= freshRanges[offset].Until {
				currentFreshRange.Until = freshRanges[offset].From - 1
			}
		}
		if currentFreshRange.From <= currentFreshRange.Until {
			count += int(currentFreshRange.Until-currentFreshRange.From) + 1
		}
	}

	return count
}

func CountFreshRanges(lines []string, freshRanges []FreshRange) int {
	var count int = 0

	for _, line := range lines {
		ingredientId := toIngredientId(line)
		if ingredientId.IsFresh(freshRanges) {
			count++
		}
	}

	return count
}

func ParseInput(blocks [][]string) []FreshRange {
	var freshRanges = make([]FreshRange, len(blocks[0]))

	for idx, line := range blocks[0] {
		parts := strings.Split(line, "-")
		freshRanges[idx] = FreshRange{
			From:  toIngredientId(parts[0]),
			Until: toIngredientId(parts[1]),
		}
	}

	return freshRanges
}

func toIngredientId(s string) IngredientId {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return IngredientId(nr)
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	freshRanges := ParseInput(blocks)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountFreshRanges(blocks[1], freshRanges))
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	freshRanges := ParseInput(blocks)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountFreshSize(freshRanges))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
