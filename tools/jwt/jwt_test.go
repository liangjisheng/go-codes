package main

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	jwt, err := GenerateToken("1234", time.Second*1)
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 2)

	claim, err := ParseToken(jwt)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claim.Address)
}
