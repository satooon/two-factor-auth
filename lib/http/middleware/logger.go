package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		url1 := c.Request.URL.String()
		method := c.Request.Method

		c.Next()

		url2 := c.Request.URL.String()
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		statusCode := strconv.Itoa(c.Writer.Status())
		comment := c.Errors.String()

		arr := []string{
			end.Format("2006/01/02 15:04:05"),
			statusCode,
			latency.String(),
			clientIP,
			method,
			url1,
			url2,
			comment,
		}
		fmt.Printf("%s\n", strings.Join(arr, "\t"))
	}
}
