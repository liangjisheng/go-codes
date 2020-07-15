package main

import (
	"encoding/base64"
)

// Base64Encode ...
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64Decode ...
func Base64Decode(src string) (string, error) {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
