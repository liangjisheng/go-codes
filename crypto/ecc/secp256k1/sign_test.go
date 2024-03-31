package secp256k1x

import "testing"

func TestSign(t *testing.T) {
	priHex, pubHex := GenerateKey()
	message := "hello world"
	signature := Signature(message, priHex)
	ok := Verify(message, pubHex, signature)
	if ok {
		t.Log("verify ok")
	} else {
		t.Log("verify not ok")
	}
}
