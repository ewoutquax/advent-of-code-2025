package day02giftshop_services

import (
	"strconv"
	"strings"

	common "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/common"
)

func ParseInput(line string) []common.RangePackageIds {
	ranges := strings.Split(line, ",")

	var rangesPackageIds []common.RangePackageIds = make([]common.RangePackageIds, len(ranges))

	for idx, raw := range ranges {
		parts := strings.Split(raw, "-")

		rangesPackageIds[idx] = common.RangePackageIds{
			From:  toPackageId(parts[0]),
			Until: toPackageId(parts[1]),
		}
	}

	return rangesPackageIds
}

func toPackageId(s string) common.PackageId {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return common.PackageId(nr)
}
