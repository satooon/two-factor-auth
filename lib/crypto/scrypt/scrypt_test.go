package scrypt_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/satooon/two-factor-auth/lib/crypto/scrypt"
)

func TestNewSalt(t *testing.T) {
	salt, err := scrypt.NewSalt()
	assert.Nil(t, err)
	assert.NotEqual(t, salt, 0)
}

func TestHash(t *testing.T) {
	salt, err := scrypt.NewSalt()
	assert.Nil(t, err)

	pass := "fdDfew2L1e3"
	h, err := scrypt.Hash(pass, salt)
	assert.Nil(t, err)
	assert.NotEqual(t, pass, h)
}
