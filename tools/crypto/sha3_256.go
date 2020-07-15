package main

import (
	"encoding/hex"

	"github.com/InWeCrypto/sha3"
)

// Sha3_256 ...
func Sha3_256(dataHexs string) (hash string, err error) {
	var dataBytes []byte
	dataBytes, err = hex.DecodeString(dataHexs)
	if err != nil {
		return "", err
	}

	digest := sha3.Sum256(dataBytes)
	return hex.EncodeToString(digest[:]), nil
}
