package helper

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	CaseSensitive:          true,
}.Froze()

// JSONEncode ...
func JSONEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONDecode ...
func JSONDecode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
