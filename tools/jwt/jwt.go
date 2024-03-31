package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Address string
	jwt.StandardClaims
}

const (
	hmacSignKey = "1234"
)

func GenerateToken(address string, expire time.Duration) (string, error) {
	// create a new token object, specifying signing method and the claims
	// you would like it to contain.
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := MyClaims{
		Address: strings.ToLower(address),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSignKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the algorithm is what your expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpectd signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret. e.g. []byte("secret")
		return []byte(hmacSignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("token claims not MyClaims type")
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}

	return claims, nil
}

func demo2() {
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
		"sub":   "example@qq.com",    // 该JWT所面向的用户
		"useID": "1000",
	})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSignKey)
	if err != nil {
		fmt.Println("sign token err:", err)
		return
	}
	fmt.Println("token:", tokenString)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the algorithm is what your expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpectd signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret. e.g. []byte("mysecret")
		return hmacSignKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("claims: %+v", claims)
		return
	}
	fmt.Println("parse err:", err)
}
