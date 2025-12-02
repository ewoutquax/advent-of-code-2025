package day02giftshop

import (
	"fmt"

	common "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/common"
	services "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/services"
	"github.com/ewoutquax/advent-of-code-2025/pkg/register"
	"github.com/ewoutquax/advent-of-code-2025/pkg/utils"
)

const Day string = "02"

func SumExtendedInvalidPackageIds(rangesPackageIds []common.RangePackageIds) int {
	var sum int = 0
	var validator = services.BuildValidator(services.WithValidationTypeExtended())

	for _, rangePackageIds := range rangesPackageIds {
		invalidIds := validator.FindInvalidPackageIds(rangePackageIds)

		for _, id := range invalidIds {
			sum += int(id)
		}
	}

	return sum
}

func SumInvalidPackageIds(rangesPackageIds []common.RangePackageIds) int {
	var sum int = 0
	var validator = services.BuildValidator()

	for _, rangePackageIds := range rangesPackageIds {
		invalidIds := validator.FindInvalidPackageIds(rangePackageIds)

		for _, id := range invalidIds {
			sum += int(id)
		}
	}

	return sum
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	ranges := services.ParseInput(line)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumInvalidPackageIds(ranges))
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	ranges := services.ParseInput(line)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumExtendedInvalidPackageIds(ranges))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
