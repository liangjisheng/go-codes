package hash

import (
	"github.com/ethereum/go-ethereum/crypto"

	"testing"
)

func TestHash(t *testing.T) {
	str := crypto.Keccak256Hash([]byte("NameRegistered(string,uint256,address)")).Hex()
	t.Log(str)
}
