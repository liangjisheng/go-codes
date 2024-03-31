package random

import (
	"fmt"
	"testing"
)

func TestRandom1(t *testing.T) {
	fmt.Println(RanNum())
}

func TestRandom2(t *testing.T) {
	fmt.Println(RanBit())
}

func TestRandom3(t *testing.T) {
	//fmt.Println(RanAToB(100, 10000))
	t.Log(GetRandomString(12))
}
