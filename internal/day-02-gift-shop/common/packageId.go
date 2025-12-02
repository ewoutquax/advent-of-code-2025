package day02giftshop_common

import "strconv"

func (rp PackageId) IsValid() bool {
	temp := strconv.Itoa(int(rp))

	if len(temp)%2 != 0 {
		return true
	}

	return !isExtendedInvalidBySize(temp, len(temp)/2)
}

func (rp PackageId) IsExtendedValid() bool {
	temp := strconv.Itoa(int(rp))

	for size := 1; size <= len(temp)/2; size++ {
		if isExtendedInvalidBySize(temp, size) {
			return false
		}
	}

	return true
}

func isExtendedInvalidBySize(pid string, size int) bool {
	if len(pid)%size != 0 {
		return false
	}

	for offset := 0; offset < len(pid); offset += size {
		if string(pid[0:size]) != string(pid[offset:offset+size]) {
			return false
		}
	}

	return true
}
