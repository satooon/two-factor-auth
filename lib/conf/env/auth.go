package env

import (
	"github.com/pquerna/otp"
)

type _auth struct {
	Issuer     string `default:"Test"`
	Period     uint   `default:"30"`
	SecretSize uint   `envconfig:"secret_size",default:"10"`
	Digits     int    `default:"6"`
	Algorithm  int    `default:"0"`
}

var auth *_auth

func Auth() *_auth {
	if auth == nil {
		auth = process("auth", &_auth{}).(*_auth)
	}
	return auth
}

func (e *_auth) GetIssuer() string {
	return e.Issuer
}

func (e *_auth) GetPeriod() uint {
	return e.Period
}

func (e *_auth) GetSecretSize() uint {
	return e.SecretSize
}

func (e *_auth) GetDigits() otp.Digits {
	return otp.Digits(e.Digits)
}

func (e *_auth) GetAlgorithm() otp.Algorithm {
	return otp.Algorithm(e.Algorithm)
}
