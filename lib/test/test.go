package test

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/satooon/two-factor-auth/asset/template"
	"github.com/satooon/two-factor-auth/lib"
	"github.com/satooon/two-factor-auth/lib/conf"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/satooon/two-factor-auth/lib/http"
	"github.com/satooon/two-factor-auth/lib/http/middleware"
)

type Engine struct {
	Engine *gin.Engine
	DB     *database.DB
}

func NewEngine() *Engine {
	if err := conf.Load(); err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()

	store, err := http.Store(env.Session())
	if err != nil {
		panic(err)
	}
	db, err := database.Open(env.Database())
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(database.Migrate...)

	if err := lib.Init(repository.NewUser(db)); err != nil {
		panic(err)
	}

	r.Use(sessions.Sessions(env.Session().Name, store))
	r.Use(middleware.Time())
	r.Use(middleware.Database(db))

	r.HTMLRender = (render.HTMLRender)(template.NewHtmlRender("layout/layout.tmpl", "layout/header.tmpl", "layout/footer.tmpl"))

	return &Engine{Engine: r, DB: db}
}
