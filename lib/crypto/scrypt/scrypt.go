package scrypt

import (
	"io"
	"encoding/base64"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"encoding/hex"
)

var (
	saltDefaultCost = 20
)

func NewSalt() (string, error) {
	b := make([]byte, saltDefaultCost)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func Hash(pass string, salt string) (string, error) {
	p := []byte(pass)
	s := []byte(salt)
	b, err := scrypt.Key(p, s, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b[:]), nil
}
