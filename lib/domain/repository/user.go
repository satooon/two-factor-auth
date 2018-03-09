package repository

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/satooon/two-factor-auth/lib/crypto/aes"
	"github.com/satooon/two-factor-auth/lib/crypto/scrypt"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
)

type User struct {
	db database.IDatabase
}

type iSecret interface {
	GetEncryptKey() string
	GetEncryptIV() string
}

type iAuth interface {
	GetIssuer() string
	GetPeriod() uint
	GetSecretSize() uint
	GetDigits() otp.Digits
	GetAlgorithm() otp.Algorithm
}

func NewUser(db database.IDatabase) *User {
	return &User{
		db: db,
	}
}

func (r *User) NewEntity(secret iSecret, auth iAuth, mail string, pass string, roleName entity.RoleName) (e *entity.User, err error) {
	m, err := aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), mail)
	if err != nil {
		return
	}
	s, err := scrypt.NewSalt()
	if err != nil {
		return
	}
	p, err := scrypt.Hash(pass, s)
	if err != nil {
		return
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      auth.GetIssuer(),
		AccountName: mail,
		Period:      auth.GetPeriod(),
		SecretSize:  auth.GetSecretSize(),
		Digits:      auth.GetDigits(),
		Algorithm:   auth.GetAlgorithm(),
	})
	if err != nil {
		return
	}
	sec, err := aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), key.Secret())
	if err != nil {
		return
	}
	e = entity.NewUser(m, s, p, sec, roleName)
	return
}

func (r *User) FindByMail(secret iSecret, mail ...string) (s entity.UserSlice, err error) {
	s = entity.UserSlice{}
	args := make([]interface{}, len(mail))
	for i := range mail {
		args[i], err = aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), mail[i])
		if err != nil {
			return
		}
	}
	err = r.db.Where("mail in (?)", args...).Find(&s).Error()
	return
}

func (r *User) Transaction(f func() error) error {
	t := r.db.Begin()
	if err := f(); err != nil {
		t.Rollback()
		return err
	}
	return t.Commit().Error()
}

func (r *User) Save(e *entity.User) error {
	if r.db.NewRecord(e) {
		return r.db.Create(e).Error()
	} else {
		return r.db.Update(e).Error()
	}
}

func (r *User) Delete(e *entity.User) error {
	return r.db.Delete(e).Error()
}
