package random

import (
	"fmt"
	"testing"
)

func TestRandom1(t *testing.T) {
	fmt.Println(Random1())
}

func TestRandom2(t *testing.T) {
	fmt.Println(Random2())
}

func TestRandom3(t *testing.T) {
	fmt.Println(Random3(100, 10000))
}
