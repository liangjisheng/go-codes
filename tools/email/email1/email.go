package email

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	//qq邮箱
	//mailConn := map[string]string{
	//	"user": "example@qq.com",
	//	"pass": "code",
	//	"host": "smtp.qq.com",
	//	"port": "465",
	//}

	//mailConn := map[string]string{
	//	"user": "example@web3.com",
	//	"pass": "xxx",
	//	"host": "smtp.gmail.com",
	//	"port": "587",
	//}

	mailConn := map[string]string{
		"user": "example@qq.com",
		"pass": "pass",
		//"host": "smtp.gmail.com",
		"host": "smtp-relay.gmail.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/plain", body)   //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
