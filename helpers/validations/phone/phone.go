package phone_validate

//
// helpers => validations => phone => phone.go
//

import (
	"bytes"
	"regexp"
	"unicode"
)

var (
	PhoneRegexp = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)?(?:((?:9\d|[2-9])\d{3})\-?(\d{4}))$`)
)

func IsValidePhone(phone string, onlyDigits, fullNationalNumber, fullInternationalNumber bool) bool {

	if !PhoneRegexp.MatchString(phone) {
		return false
	}

	phonedigits := phone
	removeNonDigits(&phonedigits)

	if fullInternationalNumber {
		if len(phonedigits) != 13 {
			return false
		}
	}

	if fullNationalNumber {
		if len(phonedigits) != 11 {
			return false
		}
	}

	if onlyDigits {
		if phone != phonedigits {
			return false
		}
	}

	return true
}

func removeNonDigits(phone *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *phone {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*phone = buf.String()
}
