package day02giftshop_services_validator

import common "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/common"

type ValidationType uint

const (
	ValidationTypeBasic ValidationType = iota + 1
	ValidationTypeExtended
)

type Validator struct {
	ValidationType
}

func (v Validator) FindInvalidPackageIds(inputRange common.RangePackageIds) []common.PackageId {
	var invalids []common.PackageId = make([]common.PackageId, 0)

	for currentPackageId := inputRange.From; currentPackageId <= inputRange.Until; currentPackageId++ {
		switch v.ValidationType {
		case ValidationTypeBasic:
			if !currentPackageId.IsValid() {
				invalids = append(invalids, currentPackageId)
			}
		case ValidationTypeExtended:
			if !currentPackageId.IsExtendedValid() {
				invalids = append(invalids, currentPackageId)
			}
		default:
			panic("No valid case found")
		}
	}

	return invalids
}
