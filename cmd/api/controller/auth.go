package controller

import (
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthController interface {
	PasswordLogin(*gin.Context)
}

type authCtl struct {
	authService service.AuthService
}

func NewAuthController(di *do.Injector) AuthController {
	return &authCtl{
		authService: do.MustInvoke[service.AuthService](di),
	}
}

func (c *authCtl) PasswordLogin(ctx *gin.Context) {
	req := &dto.PasswordLoginRequest{}
	_ = ctx.ShouldBind(req)
	resp, err := c.authService.PasswordLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
