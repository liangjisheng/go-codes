package gob

import (
	"bytes"
	"encoding/gob"
)

type Student struct {
	Name string
	Age int
	Address string
}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(src []byte, dst interface{}) error {
	buf := bytes.NewBuffer(src)
	dec := gob.NewDecoder(buf)
	return dec.Decode(dst)
}
