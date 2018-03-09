package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satooon/two-factor-auth/lib/util"
)

func Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		util.SetTime(c, &t)
		c.Next()
	}
}
