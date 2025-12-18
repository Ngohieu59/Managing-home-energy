package controller

import (
	http_response "Managing-home-energy/cmd/api/controller/http-response"
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthController interface {
	PasswordLogin(*gin.Context)
	StaffPasswordLogin(*gin.Context)
}

type authCtl struct {
	authService service.AuthService
}

func NewAuthController(di *do.Injector) AuthController {
	return &authCtl{
		authService: do.MustInvoke[service.AuthService](di),
	}
}

func (c *authCtl) Login(ctx *gin.Context, loginFunc func(context.Context, *dto.PasswordLoginRequest) (*dto.LoginResponse, error)) {
	req := &dto.PasswordLoginRequest{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	resp, err := loginFunc(ctx.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserNameNotFound):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserNotFound, err)
			return
		case errors.Is(err, dto.ErrPasswordIncorrect):
			http_response.ReturnErrMessage(ctx, http_response.ErrLoginPasswordIncorrect, err)
			return
		case errors.Is(err, dto.ErrPermissionDenied):
			http_response.ReturnErrMessage(ctx, http_response.ErrPermissionDenied, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrCannotGenerateToken, err)
			return
		}
	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}

}
func (c *authCtl) PasswordLogin(ctx *gin.Context) {
	c.Login(ctx, c.authService.PasswordLogin)
}

func (c *authCtl) StaffPasswordLogin(ctx *gin.Context) {
	c.Login(ctx, c.authService.StaffPasswordLogin)
}
