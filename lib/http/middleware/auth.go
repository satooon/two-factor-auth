package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	http2 "github.com/satooon/two-factor-auth/lib/http"
)

type iProfile interface {
	HasAuthenticated() bool
}

func Authenticate(redirectUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := http2.Session(c)
		if p, ok := s.Profile().(iProfile); ok && p.HasAuthenticated() {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, redirectUrl)
	}
}
