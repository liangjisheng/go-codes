package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func demo2() {
	hmacSmapleSecret := []byte("mysecret")

	// create a new token object, specifying signing method and the claims
	// you would like it to contain.
	issueTime := time.Now().Unix()
	t, err := time.Parse("2006-01-02 15:04:05", "2100-01-01 00:00:00")
	if err != nil {
		fmt.Println("time parse err:", err)
		return
	}
	expireTime := t.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "smallsoup",         // 该JWT的签发者
		"iat":   issueTime,           // 签发时间
		"exp":   expireTime,          // 过期时间
		"aud":   "www.smallsoup.com", // 接收该JWT的一方
		"sub":   "1294851990@qq.com", // 该JWT所面向的用户
		"useID": "1000",
	})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSmapleSecret)
	if err != nil {
		fmt.Println("sign token err:", err)
		return
	}
	fmt.Println("token:", tokenString)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the algorithm is what your expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpectd signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret. e.g. []byte("mysecret")
		return hmacSmapleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("claims: %+v", claims)
		return
	}
	fmt.Println("parse err:", err)
}
