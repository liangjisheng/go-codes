package gob

import (
	"log"
	"testing"
)

func TestGob(t *testing.T) {
	s1 := Student{
		Name:    "alice",
		Age:     18,
		Address: "shanxi",
	}

	b, err := Encode(s1)
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Printf("encode: %x\n", b)

	var s2 Student
	err = Decode(b, &s2)
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Printf("%+v\n", s2)
}
