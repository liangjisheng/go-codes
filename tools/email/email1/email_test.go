package email

import "testing"

func TestSendMail(t *testing.T) {
	//mailTo := []string{"example@qq.com", "example@web3.com", "example1@web3.com"}
	//mailTo := []string{"example@qq.com", "example@126.com"}
	mailTo := []string{"example@qq.com"}
	subject := "hello"
	body := "hello"
	err := SendMail(mailTo, subject, body)
	if err != nil {
		t.Log(err)
	}
}
