package main

import (
	crand "crypto/rand"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/coinbase/kryptology/pkg/core/curves"
	"github.com/coinbase/kryptology/pkg/sharing"
)

func getCurve(t string) *curves.Curve {
	s := strings.ToLower(t)
	if s == "ed25519" {
		return curves.ED25519()
	} else if s == "k256" {
		return curves.K256()
	} else if s == "p256" {
		return curves.P256()
	} else if s == "pallas" {
		return curves.PALLAS()
	}
	return curves.ED25519()

}

func main() {

	msg := "Hello"
	var t uint32 = uint32(3)
	var n uint32 = uint32(4)
	argCount := len(os.Args[1:])

	//curve := curves.ED25519()
	curve := getCurve("p256")

	if argCount > 0 {
		msg = os.Args[1]
	}
	if argCount > 1 {
		val, err := strconv.Atoi(os.Args[2])
		if err == nil {
			t = uint32(val)
		}
	}
	if argCount > 2 {
		val, err := strconv.Atoi(os.Args[3])
		if err == nil {
			n = uint32(val)
		}
	}
	if argCount > 3 {
		curvetype := os.Args[4]
		curve = getCurve(curvetype)
	}

	msg = "pedersen hashes yay!"
	secret := curve.NewScalar().Hash([]byte(msg))

	pt := curve.ScalarBaseMult(secret)

	pedersen, _ := sharing.NewPedersen(t, n, pt)

	shares, _ := pedersen.Split(secret, crand.Reader)

	fmt.Printf("=== Pedersen Verifiable Secret Sharing Scheme === \n")
	fmt.Printf("Curve: %s\n", curve.Name)
	fmt.Printf("Msg to hash: %x\n", msg)

	fmt.Printf("== Secret shares == %d from %d ===\n", t, n)
	for _, s := range shares.SecretShares {
		fmt.Printf("Share: %x\n", s.Bytes())
	}
	fmt.Printf("=================\n")

	fmt.Printf("== Blinding shares == %d from %d ===\n", t, n)
	for _, s := range shares.BlindingShares {
		fmt.Printf("Blinding shares: %x\n", s.Bytes())
	}
	fmt.Printf("=================\n")

	//	recovered, _ := pedersen.Combine(shares.SecretShares...)
	sG, _ := pedersen.Combine(shares.SecretShares...)
	bH, _ := pedersen.Combine(shares.BlindingShares...)

	fmt.Printf("Secret: %x\n", secret.Bytes())
	fmt.Printf("Recovered: %x\n", sG.Bytes())

	fmt.Printf("\nBlinding: %x\n", shares.Blinding.Bytes())
	fmt.Printf("Blinding Recovered: %x\n", bH.Bytes())

	err := shares.PedersenVerifier.Verify(shares.SecretShares[0], shares.BlindingShares[0])
	if err == nil {
		fmt.Printf("\nShare 1 verified\n")
	}
	err = shares.PedersenVerifier.Verify(shares.SecretShares[1], shares.BlindingShares[1])
	if err == nil {
		fmt.Printf("\nShare 2 verified\n")
	}

}
