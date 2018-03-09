package auth

import (
	"net/http"

	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/satooon/two-factor-auth/lib/conf/yaml"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
	"github.com/satooon/two-factor-auth/lib/domain/service"
	http2 "github.com/satooon/two-factor-auth/lib/http"
	"github.com/satooon/two-factor-auth/lib/http/middleware"
)

type auth struct{}

func NewHandler() *auth {
	return &auth{}
}

func (ctl *auth) Routes(r gin.IRouter) {
	g1 := r.Group("/auth")
	{
		g1.GET("", ctl.Index)
		g1.POST("/sign_in", ctl.SignIn)
		g1.POST("/sign_out", ctl.SignOut)

		g2 := g1.Group("/two_factor_auth")
		{
			g2.GET("", ctl.TwoFactorAuth)

			g3 := g2.Group("")
			{
				g3.Use(middleware.Authenticate("/auth"))
				g3.POST("/validate", ctl.TwoFactorAuthValidate)
				g3.GET("/logged_in", ctl.LoggedIn)
			}
		}
	}
}

func (ctl *auth) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/auth/index.tmpl", gin.H{
		"Error": c.QueryArray("error"),
	})
}

type SignInRequest struct {
	Mail     string `form:"mail"`
	Password string `form:"password"`
}

func (ctl *auth) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := service.NewUser(c).SignIn(req.Mail, req.Password)
	if err != nil {
		c.Redirect(http.StatusFound, "/auth?error="+url.QueryEscape("email or password illegal"))
		return
	}

	if err := http2.Session(c).SetProfile(&user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/auth/two_factor_auth")
}

func (ctl *auth) SignOut(c *gin.Context) {
	s := http2.Session(c)
	s.Clear()
	if err := s.Save(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/auth")
}

func (ctl *auth) TwoFactorAuth(c *gin.Context) {
	c.HTML(http.StatusOK, "view/auth/two_factor_auth.tmpl", gin.H{})
}

type TwoFactorAuthValidateRequest struct {
	Code int `form:"code"`
}

func (ctl *auth) TwoFactorAuthValidate(c *gin.Context) {
	var req TwoFactorAuthValidateRequest
	if err := c.Bind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, ok := http2.Session(c).Profile().(*entity.User)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err := service.NewUser(c).TwoFactorAuthCodeValidate(user, req.Code); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Redirect(http.StatusFound, "/auth/two_factor_auth/logged_in")
}

func (ctl *auth) LoggedIn(c *gin.Context) {
	if err := http2.Session(c).DeleteProxyConf(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	name := c.Query("name")
	if p, _ := yaml.Proxy().GetProxyConfSlice().First(func(p *yaml.ProxyConf) bool {
		return p.Name == name
	}); p != nil {
		if err := http2.Session(c).SetProxyConf(p); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		u := "/"
		if len(p.StripPath) > 0 {
			u = p.StripPath
		}
		c.Redirect(http.StatusFound, u)
		return
	}

	c.HTML(http.StatusOK, "view/auth/logged_in.tmpl", gin.H{
		"ProxyConfSlice": yaml.Proxy().ProxyConfSlice,
	})
}
