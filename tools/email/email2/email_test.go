package email

import "testing"

func TestSendEmail(t *testing.T) {
	user := "1294851990@qq.com"
	pass := "code"
	host := "smtp.qq.com:465"
	title := "hello"
	e := New(user, pass, host, title)

	to := []string{"liangjisheng999@126.com"}
	subject := "subject"
	body := "hello liangjisheng"

	err := e.SendMail(to, subject, body)
	if err == nil {
		t.Log("send ok")
	} else {
		t.Error(err)
		t.Log("send fail")
	}
}
