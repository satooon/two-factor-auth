package admin

import (
	"encoding/base64"
	"net/http"

	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/satooon/two-factor-auth/lib/domain/service"
	http2 "github.com/satooon/two-factor-auth/lib/http"
)

const (
	qrWidth  = 200
	qrHeight = 200

	userLimit = 20
)

type admin struct{}

func NewHandler() *admin {
	return &admin{}
}

func (ctl *admin) Routes(r gin.IRouter) {
	g1 := r.Group("/admin")
	{
		g1.GET("", ctl.Index)
		g1.POST("/sign_in", ctl.SignIn)
		g1.POST("/sign_out", ctl.SignOut)

		g2 := g1.Group("")
		{
			g2.Use(func(c *gin.Context) {
				s := http2.Session(c)
				if _, ok := s.Profile().(*entity.User); ok {
					c.Next()
					return
				}
				c.Redirect(http.StatusFound, "/admin")
			})
			g2.GET("/logged_in", ctl.LoggedIn)

			g3 := g2.Group("/two_factor_auth")
			{
				g3.GET("", ctl.TwoFactorAuth)
				g3.POST("/validate", ctl.TwoFactorAuthValidate)
			}

			g4 := g2.Group("/user")
			{
				g4.Use(func(c *gin.Context) {
					role, err := ctl.role(c)
					if err != nil {
						c.AbortWithError(http.StatusUnauthorized, err)
						return
					}
					if !role.ContainsPermission(entity.AdminUserSignUp) {
						c.AbortWithError(http.StatusUnauthorized, err)
						return
					}
				})
				g4.GET("", ctl.User)
				g4.POST("/sign_up", ctl.UserSignUp)
				g4.POST("/delete", ctl.UserDelete)
			}
		}
	}
}

func (ctl *admin) role(c *gin.Context) (*entity.Role, error) {
	user, ok := http2.Session(c).Profile().(*entity.User)
	if !ok {
		return nil, errors.New("please sign in.")
	}
	slice := repository.NewRole().FindByUser(user)
	return slice.First(func(e *entity.Role) bool { return true })
}

func (ctl *admin) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/admin/index.tmpl", gin.H{
		"Error": c.QueryArray("error"),
	})
}

type SignInRequest struct {
	Mail     string `form:"mail"`
	Password string `form:"password"`
}

func (ctl *admin) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := service.NewUser(c).SignIn(req.Mail, req.Password)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin?error="+url.QueryEscape("email or password illegal"))
		return
	}

	if err := http2.Session(c).SetProfile(&user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/admin/logged_in")
}

func (ctl *admin) SignOut(c *gin.Context) {
	s := http2.Session(c)
	s.Clear()
	if err := s.Save(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/admin")
}

func (ctl *admin) LoggedIn(c *gin.Context) {
	role, err := ctl.role(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.HTML(http.StatusOK, "view/admin/logged_in.tmpl", gin.H{
		"Role":    role,
		"Success": c.QueryArray("success"),
	})
}

func (ctl *admin) TwoFactorAuth(c *gin.Context) {
	user, ok := http2.Session(c).Profile().(*entity.User)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	key, buf, err := service.NewUser(c).QR(user, qrWidth, qrHeight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "view/admin/two_factor_auth_key.tmpl", gin.H{
		"Key":   key,
		"QR":    base64.StdEncoding.EncodeToString(buf),
		"Error": c.QueryArray("error"),
	})
}

type TwoFactorAuthValidateRequest struct {
	Code int `form:"code"`
}

func (ctl *admin) TwoFactorAuthValidate(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/admin/two_factor_auth?error="+url.QueryEscape(err.Error()))
		return
	}
	repo := repository.NewUser(database.Default(c))
	if err := repo.Transaction(func() error {
		user.SetStateTwoFactorAuth(entity.Authenticated)
		return repo.Save(user)
	}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/admin/logged_in?success="+url.QueryEscape("TwoFactorAuth Success"))
}

func (ctl *admin) User(c *gin.Context) {
	c.HTML(http.StatusOK, "view/admin/user.tmpl", gin.H{
		"Error":   c.QueryArray("error"),
		"Success": c.QueryArray("success"),
	})
}

type UserSignUpRequest struct {
	Mail     string `form:"mail"`
	Password string `form:"password"`
}

func (ctl *admin) UserSignUp(c *gin.Context) {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		c.Redirect(http.StatusFound, "/admin/user?error="+url.QueryEscape(err.Error()))
		return
	}

	repo := repository.NewUser(database.Default(c))

	user, err := repo.NewEntity(env.App(), env.Auth(), req.Mail, req.Password, entity.RoleUser)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/user?error="+url.QueryEscape(err.Error()))
		return
	}
	if err := repo.Save(user); err != nil {
type UserDeleteRequest struct {
	ID int `form:"id"`
}

func (ctl *admin) UserDelete(c *gin.Context) {
	var req UserDeleteRequest
	if err := c.Bind(&req); err != nil {
		c.Redirect(http.StatusFound, "/admin/user?error="+url.QueryEscape(err.Error()))
		return
	}

	repo := repository.NewUser(database.Default(c))

	if err := repo.Transaction(func() error {
		userSlice, err := repo.FindByID(req.ID)
		if err != nil {
			return err
		}
		user, err := userSlice.First(func(u *entity.User) bool { return true })
		if err != nil {
			return err
		}
		return repo.Delete(user)
	}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/admin/user?success="+url.QueryEscape("User delete Success"))
}
