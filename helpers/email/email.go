package email

//
// helpers => email => email.go
//

type IEmail interface {
	Initialize() error
	SetEmailCredentials(user string, password string) error
	SetSmtpServer(server string, secure bool, port int) error
	SetPopServer(server string, secure bool, port int) error
	SetImapServer(server string, secure bool, port int) error
	SendEmailFrom(address string, name string) error
	SendEmailTo(address string, name string) error
	SendEmailCc(address string, name string) error
	SendEmailBcc(address string, name string) error
	SendEmailAddHeader(key string, value ...string) error
	SendEmailSetBody(contentType string, content string) error
	SendEmailAttachFile(file string) error
	SendEmail() error
	Finally() error
}
