package main

import (
	"fmt"
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

type customPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

func demo3() {
	now := time.Now()
	pl := customPayload{
		Payload: jwt.Payload{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://baidu.com", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Foo: "foo",
		Bar: 1337,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		fmt.Println("jwt sign err")
		return
	}
	fmt.Printf("token: %+v\n", string(token))

	var pl1 customPayload
	hd, err := jwt.Verify(token, hs, &pl1)
	if err != nil {
		fmt.Println("jwt verify err:", err)
		return
	}
	fmt.Println("jwt veriry success:", hd)
}
