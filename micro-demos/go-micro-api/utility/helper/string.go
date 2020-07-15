package helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"time"
)

// MD5 ...
func MD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
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

// RandNumber ...
func RandNumber(width int) []byte {
	rand.Seed(time.Now().UnixNano())

	var buffer = bytes.NewBuffer(make([]byte, 0, width))

	for i := 0; i < width; i++ {
		_ = binary.Write(buffer, binary.BigEndian, rand.Intn(10))
	}

	return buffer.Bytes()
}
