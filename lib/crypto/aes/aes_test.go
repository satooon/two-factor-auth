package aes_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/satooon/two-factor-auth/lib/crypto/aes"
)

func TestEncrypt(t *testing.T) {
	src := "Hello World"
	key := "oedfe12340Oddq1s"

	{
		iv := "90f4d5d08ad1e85b"
		encrypt1, err := aes.Encrypt(key, iv, src)
		assert.Nil(t, err)

		encrypt2, err := aes.Encrypt(key, iv, src)
		assert.Nil(t, err)

		assert.Equal(t, encrypt1, encrypt2)
	}
	{
		encrypt1, err := aes.Encrypt(key, "90f4d5d08ad1e85b", src)
		assert.Nil(t, err)

		encrypt2, err := aes.Encrypt(key, "80f4d5d08ad1e85b", src)
		assert.Nil(t, err)

		assert.NotEqual(t, encrypt1, encrypt2)
	}
}

func TestDecrypt(t *testing.T) {
	src := "Hello World"
	key := "oedfe12340Oddq1s"
	iv := "90f4d5d08ad1e85b"

	encrypt, err := aes.Encrypt(key, iv, src)
	assert.Nil(t, err)

	decrypt, err := aes.Decrypt(key, iv, encrypt)
	assert.Nil(t, err)

	assert.Equal(t, src, decrypt)
}
