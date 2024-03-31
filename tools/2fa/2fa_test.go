package auth

import "testing"

func TestGoogleAuth(t *testing.T) {
	auth := NewGoogleAuth()
	secret := auth.GetSecret()
	t.Log("secret:", secret)
	code, _ := auth.GetCode(secret)
	t.Log("code:", code)

	qrCodeURL := auth.GetQrcodeUrl("example@qq.com", secret)
	t.Log("qrCodeURL:", qrCodeURL)

	//code = "000000"
	ok, _ := auth.VerifyCode(secret, code)
	if !ok {
		t.Error("auth verify error.")
		return
	}
	t.Log("auth verify ok.")
}
