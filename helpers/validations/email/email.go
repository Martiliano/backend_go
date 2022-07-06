package email_validate

//
// helpers => validations => email => email.go
//

// https://github.com/badoux/checkmail

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

const disconnectAfter = time.Second * 5

type SmtpServerError struct {
	Err error
}

func (e SmtpServerError) Error() string {
	return e.Err.Error()
}

func (e SmtpServerError) Code() string {
	return e.Err.Error()[0:3]
}

func NewSmtpServerError(err error) SmtpServerError {
	return SmtpServerError{
		Err: err,
	}
}

var (
	BadEmailFormatError = errors.New("Email com Formato Invalido")
	NotFoundHostError   = errors.New("Host não pode ser localizado")

	emailValidate = regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
)

func ValidateEmailFormat(email string) error {

	if strings.HasPrefix(email, " ") {
		return BadEmailFormatError
	}

	if strings.HasSuffix(email, " ") {
		return BadEmailFormatError
	}

	if !emailValidate.MatchString(strings.ToLower(email)) {
		return BadEmailFormatError
	}
	return nil
}

func ValidateEmailMX(email string) error {
	_, host := split(email)
	if _, err := net.LookupMX(host); err != nil {
		return NotFoundHostError
	}

	return nil
}

func ValidateEmailHost(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return NotFoundHostError
	}
	client, err := DialWithTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), disconnectAfter)
	if err != nil {
		return NewSmtpServerError(err)
	}
	client.Close()
	return nil
}

// https://mxtoolbox.com/SuperTool.aspx
func ValidateEmailHostAndUser(serverHostName, serverMailAddress, email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return NotFoundHostError
	}
	client, err := DialWithTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), disconnectAfter)
	if err != nil {
		return NewSmtpServerError(err)
	}
	defer client.Close()

	err = client.Hello(serverHostName)
	if err != nil {
		return NewSmtpServerError(err)
	}
	err = client.Mail(serverMailAddress)
	if err != nil {
		return NewSmtpServerError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return NewSmtpServerError(err)
	}
	return nil
}

func DialWithTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')

	if i < 0 {
		return
	}

	account = email[:i]
	host = email[i+1:]
	return
}
