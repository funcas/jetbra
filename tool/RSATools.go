package tool

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func Sign(data, pk string) (string, error) {

	decodeString, _ := base64.StdEncoding.DecodeString(pk)
	privateKey, err := x509.ParsePKCS8PrivateKey(decodeString)
	if err != nil {
		panic(err)
	}
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey),
		crypto.SHA1, hashed)
	if err != nil {

		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signature) //转换成base64返回
	return sign, nil
}

func SignBase64(originalData, privateKey string) (string, error) {
	pri := fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKey)
	block, _ := pem.Decode([]byte(pri))
	priKey, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	if parseErr != nil {
		fmt.Println(parseErr)
		return "", errors.New("解析私钥失败")
	}

	// sha256 加密方式，必须与 下面的 crypto.SHA256 对应
	// 例如使用 sha1 加密，此处应是 sha1.New()，对应 crypto.SHA1
	hash := sha1.New()
	hash.Write([]byte(originalData))
	signature, err := rsa.SignPSS(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA1, hash.Sum(nil), nil)

	return base64.StdEncoding.EncodeToString(signature), err
}
