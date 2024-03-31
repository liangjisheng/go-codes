package email

import (
	"net/smtp"
	"strings"
)

// myEmail ...
type myEmail struct {
	user     string
	password string
	host     string
	title    string
}

// New 新建一个
func New(user, password, host, title string) *myEmail {
	return &myEmail{
		user:     user,
		password: password,
		host:     host,
		title:    title,
	}
}

// SendMail 发送邮件
/*
to: 目的人
subject: 标题
body: 邮件内容
*/
func (e *myEmail) SendMail(to []string, subject, body string) error {
	err := SendToMail(e.user, e.password, e.host, e.title, subject, body, "html", to)
	if err != nil {
		return err
	}
	return nil
}

// SendToMail 发送邮件
/*
 *    user : example@example.com login smtp server user
 *    password: xxxxx login smtp server password
 *    host: smtp.example.com:port   smtp.163.com:25
 *    to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *    subject:The subject of mail
 *    body: The content of mail
 *    mailType: mail type html or text
 */
func SendToMail(user, password, host, title, subject, body, mailType string, to []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("xxj", user, password, hp[0])
	contentType := "Content-Type:text/plain;charset=UTF-8"
	if mailType == "html" {
		contentType = "Content-Type:text/html;charset=UTF-8"
	}

	msg := []byte("To:" + strings.Join(to, ";") + "\r\nFrom:" + title + "<" + user + ">\r\nSubject:" + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	return smtp.SendMail(host, auth, user, to, msg)
}
