package documents_validate

//
// helpers => validations => document => renavam.go
//

import (
	"strconv"
)

var (
	renavamAcc = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

func IsValidRENAVAM(renavam string) bool {

	if len(renavam) != 11 {
		return false
	}
	if !allDigit(renavam) {
		return false
	}

	var sum int
	for i, r := range renavam[:len(renavam)-1] {
		sum += toInt(r) * renavamAcc[i]
	}

	digit := (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}

	return string(renavam[len(renavam)-1]) == strconv.Itoa(digit)
}
