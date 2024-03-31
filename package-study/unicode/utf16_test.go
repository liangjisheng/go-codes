package unicode__test

import (
	"testing"
	"unicode/utf16"
)

func TestUTF16(t *testing.T) {
	name := []rune("世界")
	t.Log("len", len("世界"))
	t.Log("len", len(name))
	t.Log("len", len(utf16.Encode(name)))
}
