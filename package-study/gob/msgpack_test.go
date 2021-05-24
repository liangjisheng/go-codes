package gob

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"testing"
)

func TestMsgPack(t *testing.T) {
	p1 := Person{
		Name: "ljs",
		Age: 18,
		Gender: "male",
	}

	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack marshal failed,err:%v", err)
		return
	}

	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("msgpack unmarshal failed,err:%v", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
