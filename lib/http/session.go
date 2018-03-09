package http

import (
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satooon/two-factor-auth/lib/conf/env"
)

type store struct {
	sessions.Store
}

type session struct {
	sessions.Session
}

type iEnvSession interface {
	GetStore() env.SessionStore
	GetSecret() string
	GetOption() map[string]string
}

func Store(e iEnvSession) (*store, error) {
	secret := []byte(e.GetSecret())
	option := e.GetOption()
	switch e.GetStore() {
	case env.SessionStoreCookie:
		return &store{sessions.NewCookieStore(secret)}, nil
	case env.SessionStoreMemcache:
		prefix := ""
		if _, ok := option["prefix"]; ok {
			prefix = option["prefix"]
		}
		return &store{sessions.NewMemcacheStore(memcache.New(fmt.Sprintf("%s:%s", option["host"], option["port"])), prefix, secret)}, nil
	case env.SessionStoreRedis:
		size, err := strconv.Atoi(option["size"])
		if err != nil {
			return nil, err
		}
		password := ""
		if _, ok := option["password"]; ok {
			password = option["password"]
		}
		s, err := sessions.NewRedisStore(size, option["network"], fmt.Sprintf("%s:%s", option["host"], option["port"]), password, secret)
		return &store{s}, err
	}
	return nil, errors.Errorf("not match session store. %v", e.GetStore)
}

func Session(c *gin.Context) *session {
	return &session{
		Session: sessions.Default(c),
	}
}

func (s *session) SetProfile(p interface{}) error {
	s.Session.Set("profile", p)
	return s.Session.Save()
}

func (s *session) Profile() interface{} {
	return s.Session.Get("profile")
}

func (s *session) SetProxyConf(p interface{}) error {
	s.Session.Set("proxy_conf", p)
	return s.Session.Save()
}

func (s *session) ProxyConf() interface{} {
	return s.Session.Get("proxy_conf")
}

func (s *session) DeleteProxyConf() error {
	s.Session.Delete("proxy_conf")
	return s.Session.Save()
}
