package email

import "testing"

func TestSendMail(t *testing.T) {
	//mailTo := []string{"1294851990@qq.com", "zhangkai@web3.com", "liangjisheng@web3.com"}
	//mailTo := []string{"1294851990@qq.com", "liangjisheng@zks.org"}
	mailTo := []string{"liangjisheng@web3.com"}
	subject := "hello"
	body := "hello"
	err := SendMail(mailTo, subject, body)
	if err != nil {
		t.Log(err)
	}
}
