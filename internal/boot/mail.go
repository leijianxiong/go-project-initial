package boot

import (
	"crypto/tls"
	"go-project-initial/configs"
	"gopkg.in/gomail.v2"
)

var MailDialer *gomail.Dialer

func init() {
	if MailDialer == nil {
		MailDialer = gomail.NewDialer(configs.Conf.Mail.Host, configs.Conf.Mail.Port, configs.Conf.Mail.Username, configs.Conf.Mail.Password)
		MailDialer.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
}
