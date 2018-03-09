package util

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetTime(c *gin.Context, t *time.Time) {
	c.Set("time", t)
}

func Time(c *gin.Context) *time.Time {
	return c.MustGet("time").(*time.Time)
}
