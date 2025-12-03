package day03lobby

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "03"

type (
	Bank struct {
		Batteries []int
	}
)

func (b Bank) MaxUnsafeJoltage() int {
	var maxJoltages = make([]int, 12)

	for _, current := range b.Batteries {
		for idx := 0; idx < 11; idx++ {
			if maxJoltages[idx] < maxJoltages[idx+1] {
				maxJoltages[idx] = maxJoltages[idx+1]
				maxJoltages[idx+1] = 0
			}
		}
		if maxJoltages[11] < current {
			maxJoltages[11] = current
		}
	}

	var total int = 0
	for _, current := range maxJoltages {
		total = total*10 + current
	}

	return total
}

func (b Bank) MaxJoltage() int {
	var firstHigh, secondHigh int = 0, 0

	for _, current := range b.Batteries {
		if firstHigh < secondHigh {
			firstHigh = secondHigh
			secondHigh = current
		} else {
			if secondHigh < current {
				secondHigh = current
			}
		}
	}

	return firstHigh*10 + secondHigh
}

func SumMaxUnsafeJoltages(banks []Bank) int {
	var sum int = 0

	for _, bank := range banks {
		sum += bank.MaxUnsafeJoltage()
	}

	return sum
}

func SumMaxJoltages(banks []Bank) int {
	var sum int = 0

	for _, bank := range banks {
		sum += bank.MaxJoltage()
	}

	return sum
}

func ParseInput(lines []string) []Bank {
	var banks []Bank = make([]Bank, len(lines))

	for idx, line := range lines {
		var batteries = make([]int, len(line))
		for idx2, char := range strings.Split(line, "") {
			batteries[idx2] = toInt(char)
		}
		banks[idx] = Bank{
			Batteries: batteries,
		}
	}

	return banks
}

func toInt(char string) int {
	return map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}[char]
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	banks := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumMaxJoltages(banks))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	banks := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumMaxUnsafeJoltages(banks))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
