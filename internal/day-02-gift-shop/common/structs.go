package day02giftshop_common

type (
	PackageId       int
	RangePackageIds struct {
		From  PackageId
		Until PackageId
	}
)
