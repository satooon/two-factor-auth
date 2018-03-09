//go:generate gen
package entity

import (
	"time"

	"encoding/base32"
	"net/url"
	"strconv"
	"strings"

	"github.com/pquerna/otp"
	"github.com/satooon/two-factor-auth/lib/crypto/aes"
)

type StateTwoFactorAuth int

const (
	Unauthenticated StateTwoFactorAuth = 0
	Authenticated   StateTwoFactorAuth = 1
)

// +gen * slice:"Where,All,Any,First"
type User struct {
	ID                 int32  `gorm:"auto_increment;primary_key"`
	Mail               string `gorm:"unique_index; not null"`
	Salt               string `gorm:"not null"`
	Pass               string `gorm:"not null"`
	StateTwoFactorAuth StateTwoFactorAuth
	Secret             string   `gorm:"not null"`
	RoleName           RoleName `gorm:"not null;default:'user'"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}

type iSecret interface {
	GetEncryptKey() string
	GetEncryptIV() string
}

type iAuth interface {
	GetIssuer() string
	GetPeriod() uint
	GetDigits() otp.Digits
	GetAlgorithm() otp.Algorithm
}

func (u *User) TableName() string {
	return "user"
}

func NewUser(mail string, salt string, pass string, secret string, roleName RoleName) *User {
	return &User{
		Mail:     mail,
		Salt:     salt,
		Pass:     pass,
		Secret:   secret,
		RoleName: roleName,
	}
}

func (u *User) DecryptMail(s iSecret) (string, error) {
	return aes.Decrypt(s.GetEncryptKey(), s.GetEncryptIV(), u.Mail)
}

func (u *User) SetStateTwoFactorAuth(i StateTwoFactorAuth) {
	u.StateTwoFactorAuth = i
}

func (u *User) GetRoleName() RoleName {
	return RoleName(u.RoleName)
}

func (u *User) SetRoleName(n RoleName) {
	u.RoleName = n
}

func (u *User) HasAuthenticated() bool {
	return u.StateTwoFactorAuth != Unauthenticated
}

func (u *User) DecryptSecret(s iSecret) (string, error) {
	return aes.Decrypt(s.GetEncryptKey(), s.GetEncryptIV(), u.Secret)
}

// document https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func (u *User) AuthKeyUri(s iSecret, a iAuth) (*url.URL, error) {
	secret, err := u.DecryptSecret(s)
	if err != nil {
		return nil, err
	}
	mail, err := u.DecryptMail(s)
	if err != nil {
		return nil, err
	}

	// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
	v := url.Values{}
	v.Set("secret", strings.TrimRight(base32.StdEncoding.EncodeToString([]byte(secret)), "="))
	v.Set("issuer", a.GetIssuer())
	v.Set("period", strconv.FormatUint(uint64(a.GetPeriod()), 10))
	v.Set("algorithm", a.GetAlgorithm().String())
	v.Set("digits", a.GetDigits().String())

	url := &url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + a.GetIssuer() + ":" + mail,
		RawQuery: v.Encode(),
	}
	return url, nil
}

func (u *User) AuthKey(s iSecret, a iAuth) (*otp.Key, error) {
	url, err := u.AuthKeyUri(s, a)
	if err != nil {
		return nil, err
	}
	return otp.NewKeyFromURL(url.String())
}
