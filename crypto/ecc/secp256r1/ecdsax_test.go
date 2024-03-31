package ecdsax

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"testing"
)

// go test -v -run TestECDSA
func TestECDSA(t *testing.T) {
	message := []byte("hello world")

	key, err := NewSigningKey()
	if err != nil {
		return
	}
	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("privateKey", privateKey)
	publicX := hex.EncodeToString(key.X.Bytes())
	publicY := hex.EncodeToString(key.Y.Bytes())
	log.Println("publicX", publicX)
	log.Println("publicY", publicY)
	log.Println("publicKey", publicX+publicY)

	digest := sha256.Sum256(message)
	log.Println("hash", hex.EncodeToString(digest[:]))

	signature, err := Sign(message, key)
	log.Printf("signature: %x\n", signature)
	if err != nil {
		log.Println("signature fail.")
		return
	}

	if !Verify(message, signature, &key.PublicKey) {
		log.Println("verify fail.")
		return
	}
	log.Println("verify success.")
}
