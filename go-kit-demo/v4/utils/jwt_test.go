package utils

import (
	"fmt"
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	jwtToken, err := CreateJwtToken("alice", 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(jwtToken)
	jwtInfo, err := ParseToken(jwtToken)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(jwtInfo)
}
