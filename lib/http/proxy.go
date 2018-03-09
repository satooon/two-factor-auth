package http

import (
	"net/http/httputil"

	"net/http"

	"regexp"

	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/conf/yaml"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
)

type _proxy struct {
	conf iConfProxy
}

type iConfProxy interface {
	GetProxyConfSlice() yaml.ProxyConfSlice
}

const (
	headerLocation = "Location"
	headerAuthUser = "X-Auth-User-Mail"
	staticFile     = `\.(jpg|JPG|gif|GIF|png|PNG|swf|SWF|css|CSS|js|JS|inc|INC|ico|ICO)$`
)

func NewProxy(c iConfProxy) *_proxy {
	return &_proxy{
		conf: c,
	}
}

func (m *_proxy) Proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := Session(c)
		p, ok := s.ProxyConf().(*yaml.ProxyConf)
		if !ok {
			c.Next()
			return
		}
		user, ok := s.Profile().(*entity.User)
		if !ok {
			c.Next()
			return
		}
		d, err := p.DestUrl()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		mail, err := user.DecryptMail(env.App())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = d.Scheme
			req.URL.Host = d.Host
			req.URL.Path = d.Path + req.URL.Path
			req.Header.Set(headerAuthUser, mail)
			c.Request = req
			c.Request.Host = d.Host
		}}
		proxy.ServeHTTP(c.Writer, c.Request)

		if u := c.Writer.Header().Get(headerLocation); len(u) > 0 {
			d, err := url.Parse(u)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.Writer.Header().Del(headerLocation)
			c.Redirect(http.StatusFound, d.RequestURI())
			return
		}
		c.Abort()
	}
}

func (m *_proxy) Static() gin.HandlerFunc {
	reg := regexp.MustCompile(staticFile)
	return func(c *gin.Context) {
		if !reg.MatchString(c.Request.URL.String()) {
			c.Next()
			return
		}
		p, ok := Session(c).ProxyConf().(*yaml.ProxyConf)
		if !ok {
			c.Next()
			return
		}
		url, err := p.DestUrl()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			c.Request = req
		}}
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
