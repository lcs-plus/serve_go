package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	hashedBytes := hash.Sum(nil)
	hashedString := fmt.Sprintf("%x", hashedBytes)
	return hashedString
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("解码出错:", err)
		return "", err
	}
	decodedString := string(decodedBytes)
	return decodedString, nil
}
