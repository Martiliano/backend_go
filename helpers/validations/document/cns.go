package documents_validate

//
// helpers => validations => document => cns.go
//

import (
	"regexp"
)

var (
	CNSRegexp = regexp.MustCompile(`^([12]\d{2}\s?\d{4}\s?\d{4}\s?00[01]\d|[789]\d{2}\s?\d{4}\s?\d{4}\s?\d{4})$`)
)

func IsValidCNS(doc string) bool {

	if !CNSRegexp.MatchString(doc) {
		return false
	}

	removeNonDigits(&doc)

	sum := 0
	for i, r := range doc {
		sum += toInt(r) * (15 - i)
	}

	return sum%11 == 0
}
