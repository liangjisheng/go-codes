package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func demo1() {
	head1 := `{"typ":"JWT","alg":"HS256"}`
	head1Base64 := base64.StdEncoding.EncodeToString([]byte(head1))

	payload1 := `{"iss":"smallsoup","iat":1528902195,"exp":1528988638,"aud":"www.smallsoup.com","sub":"smallsoup@qq.com","userId":"0418"}`
	payload1Base64 := base64.StdEncoding.EncodeToString([]byte(payload1))

	encodedstring := head1Base64 + "." + payload1Base64

	hash := hmac.New(sha256.New, []byte("mysecret"))
	hash.Write([]byte(encodedstring))

	signature := strings.TrimRight(base64.URLEncoding.EncodeToString(hash.Sum(nil)), "=")
	fmt.Println("signature:", signature)

	token := encodedstring + "." + signature
	fmt.Println("token:", token)
}
