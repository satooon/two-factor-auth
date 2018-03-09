package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satooon/two-factor-auth/lib/controller/admin"
	"github.com/satooon/two-factor-auth/lib/controller/auth"
)

func Routes(r gin.IRouter) {
	auth.NewHandler().Routes(r)
}

func AdminRoutes(r gin.IRouter) {
	admin.NewHandler().Routes(r)
}
