package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
)

func eccSign() error {
	message := []byte("Hello world")

	key, err := NewSigningKey()
	if err != nil {
		return err
	}

	signature, err := Sign(message, key)
	if err != nil {
		return err
	}
	fmt.Printf("after sign: %x\n", signature)

	if !Verify(message, signature, &key.PublicKey) {
		fmt.Printf("验证失败")
		return errors.New("verify failed")
	}

	fmt.Printf("验证成功.")
	return nil
}

// NewSigningKey ...
func NewSigningKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return key, err
}

// Sign ...
func Sign(data []byte, privkey *ecdsa.PrivateKey) ([]byte, error) {
	// hash message
	digest := sha256.Sum256(data)

	// sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, privkey, digest[:])
	if err != nil {
		return nil, err
	}

	// encode the signature {r, s}
	// big.Int.Bytes() will need padding the case of leading zero bytes.
	params := privkey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// Verify checks a raw ECDSA signature
// Returns true if it's valid and false if not.
func Verify(data, signature []byte, pubkey *ecdsa.PublicKey) bool {
	digest := sha256.Sum256(data)
	curveOrderByteSize := pubkey.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])

	return ecdsa.Verify(pubkey, digest[:], r, s)
}
