// goCrypto project goCrypto.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//原始数据
	var data = []byte("fang")
	//md5
	hash := md5.New()
	hash.Write(data)
	encryptedData := hash.Sum(nil)
	fmt.Println("md5:", hex.EncodeToString(encryptedData))

	//rsa加密
	encryptData, err := RsaEncrypt(data, PublicKey)
	if err != nil {
		fmt.Println("encrypt err:", err)
	}
	fmt.Println("encryptData:", encryptData)
	//rsa解密
	decryptData, err := RsaDecrypt(encryptData, PrivateKey)
	if err != nil {
		fmt.Println("decrypt err:", err)
	}
	fmt.Println("decryptData:", decryptData)
	//加签
	signData, err := RsaSignWithSha1Base64(data, PrivateKey)
	if err != nil {
		fmt.Println("sign err:", err)
	}
	fmt.Println("signData:", signData)
	//验签
	verifyResult := RsaVerySignWithSha1Base64(signData, data, PublicKey)
	if verifyResult != nil {
		fmt.Println("verify failed:", verifyResult)
	}
	fmt.Println("verifyResult:", verifyResult)
}
