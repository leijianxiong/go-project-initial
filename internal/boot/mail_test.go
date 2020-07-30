package boot

import (
	"fmt"
	"go-project-initial/configs"
	"gopkg.in/gomail.v2"
	"log"
	"testing"
	"time"
)

func TestMailDialer(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", configs.Conf.Mail.Username)
	m.SetHeader("To", "jianxiong.lei@outlook.com")
	m.SetHeader("Subject", fmt.Sprintf("TestMailDialer-%s", time.Now()))
	m.SetBody("text/html", fmt.Sprintf("<p>message: %s</p>", time.Now()))
	err := MailDialer.DialAndSend(m)
	log.Println("send mail err:", err)
}
