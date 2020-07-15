package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/ripemd160"
)

// MD5 ...
func MD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

// Ripemd160 ...
func Ripemd160(str string) string {
	_ripemd160 := ripemd160.New()
	_ripemd160.Write([]byte(str))
	return hex.EncodeToString(_ripemd160.Sum([]byte(nil)))
}

// MD5FILE ...
func MD5FILE(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	_md5 := md5.New()
	_, _ = io.Copy(_md5, f)
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

// SHA1 ...
func SHA1(str string) string {
	_sha1 := sha1.New()
	_sha1.Write([]byte(str))
	return hex.EncodeToString(_sha1.Sum([]byte(nil)))
}

// SHA256 ...
func SHA256(str string) string {
	_sha256 := sha256.New()
	_sha256.Write([]byte(str))
	return hex.EncodeToString(_sha256.Sum([]byte(nil)))
}

// HMAC ...
func HMAC(key, data string) string {
	_hmac := hmac.New(md5.New, []byte(key))
	_hmac.Write([]byte(data))
	return hex.EncodeToString(_hmac.Sum([]byte(nil)))
}
