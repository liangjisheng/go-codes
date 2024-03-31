package secp256k1x

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec"
)

func GenerateKey() (priKeyHex, pubKeyHex string) {
	curve := btcec.S256()
	priKey, err := btcec.NewPrivateKey(curve)
	if err != nil {
		return "", ""
	}
	priKeyHex = hex.EncodeToString(priKey.Serialize())

	pubKey := priKey.PubKey()
	pubKeyHex = hex.EncodeToString(pubKey.SerializeUncompressed())
	return priKeyHex, pubKeyHex
}

func Signature(message, priKeyHex string) string {
	privateKeyBytes, _ := hex.DecodeString(priKeyHex)
	curve := btcec.S256()
	pri, _ := btcec.PrivKeyFromBytes(curve, privateKeyBytes)

	digest := sha256.Sum256([]byte(message))
	signature, err := pri.Sign(digest[:])
	if err != nil {
		return ""
	}
	signHex := hex.EncodeToString(signature.Serialize())
	return signHex
}

func Verify(message, pubKeyHex, signature string) bool {
	digest := sha256.Sum256([]byte(message))

	curve := btcec.S256()
	publicKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return false
	}
	pub, err := btcec.ParsePubKey(publicKeyBytes, curve)
	if err != nil {
		return false
	}

	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	signParsed, err := btcec.ParseSignature(signatureBytes, curve)
	if err != nil {
		return false
	}

	return signParsed.Verify(digest[:], pub)
}

func demo() {
	privateKey := "117ab6a0ba9d19ce285ae42607c48abce125abaffc42d3f9926af6e9c16b833c"
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	//publicKey := "04944a664fd9b0447350ee2f022123285c2f0db8bc0738046aeb1fe712c309e989bce802209809252d1b44c83abd83a544fbc2603b02472c38a5689c094e4ad059"
	publicKey := "03944a664fd9b0447350ee2f022123285c2f0db8bc0738046aeb1fe712c309e989"
	publicKeyBytes, _ := hex.DecodeString(publicKey)

	curve := btcec.S256()

	pub, err := btcec.ParsePubKey(publicKeyBytes, curve)
	if err != nil {
		log.Println(err)
		return
	}

	pubCompressedHex := hex.EncodeToString(pub.SerializeCompressed())
	log.Println("pubCompressedHex", pubCompressedHex)
	//03944a664fd9b0447350ee2f022123285c2f0db8bc0738046aeb1fe712c309e989

	pubHex := hex.EncodeToString(pub.SerializeUncompressed())
	if pubHex == publicKey {
		log.Println("public ok")
	}

	pri, _ := btcec.PrivKeyFromBytes(curve, privateKeyBytes)
	priHex := hex.EncodeToString(pri.Serialize())
	if priHex == privateKey {
		log.Println("private ok")
	}

	message := []byte("hello world")
	digest := sha256.Sum256(message)
	digestHex := hex.EncodeToString(digest[:])
	log.Println("digestHex", digestHex)
	signature, err := pri.Sign(digest[:])
	if err != nil {
		log.Println(err)
		return
	}
	signHex := hex.EncodeToString(signature.Serialize())
	log.Println("sign", signHex)

	sign1 := "3044022059fb47ee486989b0448a7fb7ec7a063c6c8ad04df66871ede14279547df27a1d02204f01867d58885460b772809212e516f320377b474b993c14c037153984e75037"
	sign1Bytes, _ := hex.DecodeString(sign1)
	signParse, err := btcec.ParseSignature(sign1Bytes, curve)
	if err != nil {
		log.Println(err)
		return
	}
	signHex1 := hex.EncodeToString(signParse.Serialize())
	if signHex1 == signHex {
		log.Println("signature ok")
	}

	ok := signParse.Verify(digest[:], pub)
	if ok {
		log.Println("verify ok")
	}
}
