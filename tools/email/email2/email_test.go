package email

import "testing"

func TestSendEmail(t *testing.T) {
	user := "example@qq.com"
	pass := "code"
	host := "smtp.qq.com:465"
	title := "hello"
	e := New(user, pass, host, title)

	to := []string{"example@126.com"}
	subject := "subject"
	body := "hello alice"

	err := e.SendMail(to, subject, body)
	if err == nil {
		t.Log("send ok")
	} else {
		t.Error(err)
		t.Log("send fail")
	}
}
