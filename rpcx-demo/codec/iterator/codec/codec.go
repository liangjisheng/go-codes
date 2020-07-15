package codec

import (
	jsoniter "github.com/json-iterator/go"
)

// JsoniterCodec ...
type JsoniterCodec struct {
}

// Decode ...
func (c *JsoniterCodec) Decode(data []byte, i interface{}) error {
	return jsoniter.Unmarshal(data, i)
}

// Encode ...
func (c *JsoniterCodec) Encode(i interface{}) ([]byte, error) {
	return jsoniter.Marshal(i)
}
