package main

import (
	"fmt"

	http2 "net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/labstack/gommon/log"
	"github.com/satooon/two-factor-auth/asset/template"
	"github.com/satooon/two-factor-auth/lib"
	"github.com/satooon/two-factor-auth/lib/conf"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/conf/version"
	"github.com/satooon/two-factor-auth/lib/controller"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/satooon/two-factor-auth/lib/http"
	"github.com/satooon/two-factor-auth/lib/http/middleware"
)

func main() {
	log.Printf("Version: %s", version.Version)

	if err := conf.Load(); err != nil {
		panic(err)
	}

	r := gin.New()

	store, err := http.Store(env.Session())
	if err != nil {
		panic(err)
	}
	db, err := database.Open(env.Database())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := lib.Init(repository.NewUser(db)); err != nil {
		panic(err)
	}

	r.Use(sessions.Sessions(env.Session().Name, store))
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Time())
	r.Use(middleware.Database(db))

	r.HTMLRender = (render.HTMLRender)(template.NewHtmlRender("layout/layout.tmpl", "layout/header.tmpl", "layout/footer.tmpl"))

	controller.AdminRoutes(r)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http2.StatusFound, "/admin")
	})

	log.Fatal(r.Run(fmt.Sprintf(":%v", env.App().GetAdminPort())))
}
