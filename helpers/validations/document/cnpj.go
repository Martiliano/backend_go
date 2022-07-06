package documents_validate

//
// helpers => validations => document => cnpj.go
//

// https://github.com/paemuri/brdoc

import (
	"regexp"
)

var (
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

func IsValidCNPJ(cnpj string, onlyDigits bool) bool {

	const (
		size = 12
		pos  = 5
	)

	isValid := isValidCnpj(cnpj, CNPJRegexp, size, pos)

	if isValid && onlyDigits {
		cnpjOnlyDigits := cnpj
		removeNonDigits(&cnpjOnlyDigits)

		if cnpj != cnpjOnlyDigits {
			isValid = false
		}
	}

	return isValid
}

func isValidCnpj(doc string, pattern *regexp.Regexp, size int, position int) bool {

	if !pattern.MatchString(doc) {
		return false
	}

	removeNonDigits(&doc)

	if allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := calcVerifyDigit(d, position)

	d = d + digit
	digit = calcVerifyDigit(d, position+1)

	return doc == d+digit
}
