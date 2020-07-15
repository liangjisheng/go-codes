package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	src := "你好，世界！ hello world"
	fmt.Println(src)
	dst := Base64Encode(src)
	fmt.Println(dst)

	srcDecode, err := Base64Decode(dst)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(srcDecode)
}

func TestKeccak256(t *testing.T) {
	dataHexs := "123456"
	hash, err := Keccak256(dataHexs)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("hash:", hash)
}

func TestSha3_256(t *testing.T) {
	dataHexs := "123456"
	hash, err := Sha3_256(dataHexs)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("hash:", hash)
}

func TestAesCFBCrypto(t *testing.T) {
	plaintext := []byte("My name is Astaxie")
	key := []byte("astaxie12798akljzmknm.ahkjkljl;k")
	dst, err := AesCFBCrypto(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("dst:", dst)

	plain, err := AesCFBDecrypto(dst, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(plain))
}

func TestAesCBCEncrypt(t *testing.T) {
	// 私钥长度只支持128　192　256位的任意字节数组
	var aeskey = []byte("这就是私钥,自己定义呦.")
	fmt.Println("私钥：", string(aeskey))
	fmt.Printf("私钥长度:%d字节\n", len(aeskey))
	pass := []byte("罗小黑战记大电影今天上映了，真好看，快去看啊")
	fmt.Println("原文：", string(pass))
	xpass, err := AesCBCEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("==============进行加密==============")
	fmt.Printf("加密信息:%x\n", xpass)
	fmt.Println("===================================")
	tpass, err := AesCBCDecrypt(xpass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("==============进行解密==============")
	fmt.Printf("解密后的原文:%s\n", tpass)
	fmt.Println("===================================")
}

func TestRSA(t *testing.T) {
	// rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	bits := 1024
	prvKey, pubKey := GenRsaKey(bits)
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData := RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := RsaEncrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}

func TestGetRand(t *testing.T) {
	GetRand()
}
