package documents_validate

//
// helpers => validations => document => cpf.go
//

// https://github.com/paemuri/brdoc

import (
	"regexp"
)

var (
	CPFRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
)

func IsValidCPF(cpf string, onlyDigits bool) bool {

	const (
		size = 9
		pos  = 10
	)

	isValid := isValidCpf(cpf, CPFRegexp, size, pos)

	if isValid && onlyDigits {
		cpfOnlyDigits := cpf
		removeNonDigits(&cpfOnlyDigits)

		if cpf != cpfOnlyDigits {
			isValid = false
		}
	}

	return isValid
}

func isValidCpf(doc string, pattern *regexp.Regexp, size int, position int) bool {

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
