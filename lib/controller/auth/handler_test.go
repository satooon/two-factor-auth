package auth_test

import (
	"net/http"
	"testing"

	"time"

	"os"

	"github.com/appleboy/gofight"
	"github.com/pquerna/otp/totp"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/controller/auth"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/satooon/two-factor-auth/lib/test"
	"github.com/stretchr/testify/assert"
)

var (
	engine *test.Engine
)

func TestMain(m *testing.M) {
	engine = test.NewEngine()
	defer engine.DB.Close()

	auth.NewHandler().Routes(engine.Engine)

	os.Exit(m.Run())
}

func TestAuth_Index(t *testing.T) {
	r := gofight.New()
	r.GET("/auth").SetDebug(true).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, res.Code, http.StatusOK)
		assert.Contains(t, res.Body.String(), "SignIn")
	})
}

func TestAuth_SignIn(t *testing.T) {
	r := gofight.New()

	r.POST("/auth/sign_in").SetForm(gofight.H{
		"mail":     env.App().GetAdminMail(),
		"password": env.App().GetAdminPass(),
	}).SetDebug(true).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, res.Code, http.StatusFound)
	})
}

func TestAuth_SignOut(t *testing.T) {
	repo := repository.NewUser(engine.DB)
	us, err := repo.FindByMail(env.App(), env.App().GetAdminMail())
	assert.Nil(t, err)
	assert.Len(t, us, 1)
	us[0].SetStateTwoFactorAuth(entity.Unauthenticated)
	assert.Nil(t, repo.Save(us[0]))

	r := gofight.New().SetDebug(true)
	r.POST("/auth/sign_in").SetForm(gofight.H{
		"mail":     env.App().GetAdminMail(),
		"password": env.App().GetAdminPass(),
	}).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, res.Code, http.StatusFound)

		r := gofight.New().SetDebug(true)
		cookie := gofight.H{}
		for _, v := range (*res).Result().Cookies() {
			cookie[v.Name] = v.Value
		}
		r.POST("/auth/sign_out").SetCookie(cookie).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusFound)
		})
	})
}

func TestAuth_TwoFactorAuth(t *testing.T) {
	repo := repository.NewUser(engine.DB)
	us, err := repo.FindByMail(env.App(), env.App().GetAdminMail())
	assert.Nil(t, err)
	assert.Len(t, us, 1)
	us[0].SetStateTwoFactorAuth(entity.Unauthenticated)
	assert.Nil(t, repo.Save(us[0]))

	r := gofight.New().SetDebug(true)
	r.POST("/auth/sign_in").SetForm(gofight.H{
		"mail":     env.App().GetAdminMail(),
		"password": env.App().GetAdminPass(),
	}).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, res.Code, http.StatusFound)

		r := gofight.New().SetDebug(true)
		cookie := gofight.H{}
		for _, v := range (*res).Result().Cookies() {
			cookie[v.Name] = v.Value
		}
		r.GET("/auth/two_factor_auth").SetCookie(cookie).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusOK)
		})
	})
}

func TestAuth_TwoFactorAuthValidate(t *testing.T) {
	repo := repository.NewUser(engine.DB)
	us, err := repo.FindByMail(env.App(), env.App().GetAdminMail())
	assert.Nil(t, err)
	assert.Len(t, us, 1)
	us[0].SetStateTwoFactorAuth(entity.Unauthenticated)
	assert.Nil(t, repo.Save(us[0]))

	r := gofight.New().SetDebug(true)
	r.POST("/auth/sign_in").SetForm(gofight.H{
		"mail":     env.App().GetAdminMail(),
		"password": env.App().GetAdminPass(),
	}).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, res.Code, http.StatusFound)

		r := gofight.New().SetDebug(true)
		cookie := gofight.H{}
		for _, v := range (*res).Result().Cookies() {
			cookie[v.Name] = v.Value
		}
		key, err := us[0].AuthKey(env.App(), env.Auth())
		assert.Nil(t, err)
		code, err := totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{
			Period:    env.Auth().GetPeriod(),
			Skew:      1,
			Digits:    env.Auth().GetDigits(),
			Algorithm: env.Auth().GetAlgorithm(),
		})
		assert.Nil(t, err)

		r.POST("/auth/two_factor_auth/validate").SetForm(gofight.H{
			"code": code,
		}).SetCookie(cookie).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusFound)
		})

		r.POST("/auth/two_factor_auth/validate").SetForm(gofight.H{
			"code": "123456",
		}).SetCookie(cookie).Run(engine.Engine, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, res.Code, http.StatusUnauthorized)
		})
	})
}
