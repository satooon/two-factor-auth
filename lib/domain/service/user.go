package service

import (
	"fmt"

	"bytes"

	"image/png"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/crypto/scrypt"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/satooon/two-factor-auth/lib/util"
)

type user struct {
	c  *gin.Context
	db *database.DB
}

func NewUser(c *gin.Context) *user {
	return &user{
		c:  c,
		db: database.Default(c),
	}
}

func (s *user) SignIn(mail, pass string) (*entity.User, error) {
	repo := repository.NewUser(s.db)
	us, err := repo.FindByMail(env.App(), mail)
	if err != nil {
		return nil, err
	}
	if len(us) <= 0 {
		return nil, fmt.Errorf("not found. mail:%s", mail)
	}
	p, err := scrypt.Hash(pass, us[0].Salt)
	if err != nil {
		return nil, err
	}
	if us[0].Pass != p {
		return nil, fmt.Errorf("password not match. pass:%s", pass)
	}
	return us[0], nil
}

func (s *user) QR(u *entity.User, width, height int) (*otp.Key, []byte, error) {
	key, err := u.AuthKey(env.App(), env.Auth())
	if err != nil {
		return nil, nil, err
	}
	img, err := key.Image(width, height)
	if err != nil {
		return nil, nil, errors.New("key image failed")
	}
	var buf bytes.Buffer
	png.Encode(&buf, img)

	return key, buf.Bytes(), nil
}

func (s *user) TwoFactorAuthCodeValidate(u *entity.User, code int) error {
	key, err := u.AuthKey(env.App(), env.Auth())
	if err != nil {
		return err
	}
	valid, err := totp.ValidateCustom(strconv.Itoa(code), key.Secret(), *util.Time(s.c), totp.ValidateOpts{
		Period:    env.Auth().GetPeriod(),
		Skew:      1,
		Digits:    env.Auth().GetDigits(),
		Algorithm: env.Auth().GetAlgorithm(),
	})
	if !valid || err != nil {
		if err == nil {
			err = errors.New("code mismatch")
		}
		return err
	}
	return nil
}
