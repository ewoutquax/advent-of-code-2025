package day02giftshop_test

import (
	"strings"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop"
	common "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/common"
	services "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/services"
	"github.com/stretchr/testify/assert"
)

func TestFindInvalidPackageIds(t *testing.T) {
	testCases := map[common.RangePackageIds][]common.PackageId{
		{From: 11, Until: 22}:                 {11, 22},
		{From: 95, Until: 115}:                {99},
		{From: 998, Until: 1012}:              {1010},
		{From: 1188511880, Until: 1188511890}: {1188511885},
		{From: 222220, Until: 222224}:         {222222},
		{From: 1698522, Until: 1698528}:       {},
		{From: 446443, Until: 446449}:         {446446},
		{From: 38593856, Until: 38593862}:     {38593859},
	}

	validator := services.BuildValidator()
	for inputRange, expectedPackageIds := range testCases {
		actualPackageIds := validator.FindInvalidPackageIds(inputRange)
		assert.Equal(t, expectedPackageIds, actualPackageIds)
	}
}

func TestFindExtendedInvalidPackageIds(t *testing.T) {
	testCases := map[common.RangePackageIds][]common.PackageId{
		{From: 11, Until: 22}:                 {11, 22},
		{From: 95, Until: 115}:                {99, 111},
		{From: 998, Until: 1012}:              {999, 1010},
		{From: 1188511880, Until: 1188511890}: {1188511885},
		{From: 222220, Until: 222224}:         {222222},
		{From: 1698522, Until: 1698528}:       {},
		{From: 446443, Until: 446449}:         {446446},
		{From: 38593856, Until: 38593862}:     {38593859},
		{From: 565653, Until: 565659}:         {565656},
		{From: 824824821, Until: 824824827}:   {824824824},
		{From: 2121212118, Until: 2121212124}: {2121212121},
	}

	validator := services.BuildValidator(services.WithValidationTypeExtended())
	for inputRange, expectedPackageIds := range testCases {
		actualPackageIds := validator.FindInvalidPackageIds(inputRange)
		assert.Equal(t, expectedPackageIds, actualPackageIds)
	}
}

func TestSumInvalidPackageIds(t *testing.T) {
	rangesPackagesIds := services.ParseInput(testInput())

	sum := SumInvalidPackageIds(rangesPackagesIds)

	assert.Equal(t, 1227775554, sum)
}

func TestSumExtendedInvalidPackageIds(t *testing.T) {
	rangesPackagesIds := services.ParseInput(testInput())

	sum := SumExtendedInvalidPackageIds(rangesPackagesIds)

	assert.Equal(t, 4174379265, sum)
}

func testInput() string {
	return strings.Join(
		[]string{
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,",
			"1698522-1698528,446443-446449,38593856-38593862,565653-565659,",
			"824824821-824824827,2121212118-2121212124",
		}, "")
}
