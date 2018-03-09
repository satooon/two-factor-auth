package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satooon/two-factor-auth/lib/database"
)

func Database(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		database.SetContext(c, db)
		c.Next()
	}
}
