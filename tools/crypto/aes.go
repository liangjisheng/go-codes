package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// AesCFBCrypto ...
func AesCFBCrypto(plain, key []byte) (dst []byte, err error) {
	var c cipher.Block
	c, err = aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plain))
	cfb.XORKeyStream(ciphertext, []byte(plain))
	return ciphertext, nil
}

// AesCFBDecrypto ...
func AesCFBDecrypto(src, key []byte) (dst []byte, err error) {
	var c cipher.Block
	c, err = aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBDecrypter(c, commonIV)
	plain := make([]byte, len(src))
	cfb.XORKeyStream(plain, []byte(src))
	return plain, nil
}

// PKCS5Padding ...
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding ...
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesCBCEncrypt ...
func AesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesCBCDecrypt ...
func AesCBCDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
