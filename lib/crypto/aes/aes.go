package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
	"bytes"
	"github.com/pkg/errors"
)

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}
	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func UnPad(src []byte) ([]byte, error) {
	length := len(src)
	unPadding := int(src[length-1])
	if unPadding > length {
		return nil, errors.New("UnPad error. This could happen when incorrect encryption key is used")
	}
	return src[:(length - unPadding)], nil
}

func Encrypt(key string, iv string, text string) (string, error) {
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	msg := Pad([]byte(text))
	cipherText := make([]byte, aes.BlockSize+len(msg))
	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(cipherText))
	return finalMsg, nil
}

func Decrypt(key string, iv string, text string) (string, error) {
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}
	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("BlockSize must be multiple of decoded message length")
	}
	msg := decodedMsg[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, []byte(iv))
	cfb.XORKeyStream(msg, msg)
	unPadMsg, err := UnPad(msg)
	if err != nil {
		return "", err
	}
	return string(unPadMsg), nil
}

