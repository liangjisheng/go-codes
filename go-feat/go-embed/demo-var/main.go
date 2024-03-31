package main

import (
	_ "embed"
	"fmt"
)

//go:embed version.txt
var version string

//go:embed version.txt
var versionByte []byte

func main() {
	fmt.Printf("version: %q\n", version)
	fmt.Printf("versionByte: %q\n", string(versionByte))
}
