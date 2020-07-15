package main

import (
	"encoding/hex"

	"github.com/InWeCrypto/sha3"
)

// Keccak256 ...
func Keccak256(dataHexs string) (hash string, err error) {
	keccak256 := sha3.NewKeccak256()
	var dataBytes []byte
	dataBytes, err = hex.DecodeString(dataHexs)
	if err != nil {
		return "", err
	}

	_, err = keccak256.Write(dataBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(keccak256.Sum([]byte(nil))), nil
}
