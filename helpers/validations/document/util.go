package documents_validate

//
// helpers => validations => document => util.go
//

import (
	"bytes"
	"strconv"
	"unicode"
)

func allDigit(doc string) bool {

	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func toInt(r rune) int {
	return int(r - '0')
}

func removeNonDigits(doc *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func allEq(cpf string) bool {

	base := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if base != cpf[i] {
			return false
		}
	}

	return true
}

func calcVerifyDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
