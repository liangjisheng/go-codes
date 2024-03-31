package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
)

func encode() {
	// Encode example data with the modified base58 encoding scheme.
	data := []byte("Test data")
	encoded := base58.Encode(data)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)
}

func decode() {
	// Decode example modified base58 encoded data.
	encoded := "25JnwSn7XKfNQ"
	decoded := base58.Decode(encoded)

	// Show the decoded data.
	fmt.Println("Decoded Data:", string(decoded))
}

func checkEncode() {
	// Encode example data with the Base58Check encoding scheme.
	data := []byte("Test data")
	encoded := base58.CheckEncode(data, 0)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)
}

func checkDecode() {
	// Decode an example Base58Check encoded data.
	//encoded := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	encoded := "182iP79GRURMp7oMHDU"
	decoded, version, err := base58.CheckDecode(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show the decoded data.
	fmt.Printf("Decoded data: %s\n", string(decoded))
	fmt.Printf("Decoded data: %x\n", decoded)
	fmt.Println("Version Byte:", version)
}

func main() {
	//encode()
	//decode()

	//checkEncode()
	checkDecode()
}
