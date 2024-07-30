package main

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestSha256(t *testing.T) {
	resBytes := SHA256([]byte("hello world"))
	resHex := hex.EncodeToString(resBytes)
	log.Println(resHex)

	resBytes1 := SHA256([]byte("hello world"))
	resHex1 := hex.EncodeToString(resBytes1)
	log.Println(resHex1)

	// Output
	// b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
	// b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}

func TestKeccak256(t *testing.T) {
	str := crypto.Keccak256Hash([]byte("NameRegistered(string,uint256,address)")).Hex()
	t.Log(str)
}
