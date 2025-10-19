package controller

import (
	http_response "Managing-home-energy/cmd/api/controller/http-response"
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type UserController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

type userCtl struct {
	userService service.UserService
}

func NewUserController(di *do.Injector) UserController {
	return &userCtl{
		userService: do.MustInvoke[service.UserService](di),
	}
}

func (uc *userCtl) Create(ctx *gin.Context) {
	req := &dto.CreateUserReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	resp, err := uc.userService.CreateUser(ctx, req)
	switch {
	case err == nil:
		break
	case errors.Is(err, dto.ErrUserIDNotFound):
		http_response.ReturnErrMessage(ctx, http_response.ErrUserNotFound, err)
		return
	case errors.Is(err, dto.ErrUserAlreadyExists):
		http_response.ReturnErrMessage(ctx, http_response.ErrUserAlreadyExists, errors.New("cannot create same username"))
		return
	case errors.Is(err, dto.ErrNotFillAllFields):
		http_response.ReturnErrMessage(ctx, http_response.ErrNotFillAllFields, errors.New("some fields are missing"))
		return
	case errors.Is(err, dto.ErrDataFormatWrong):
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, err)
		return
	default:
		http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
		return
	}
	http_response.ReturnSuccessMessage(ctx, resp)
}

func (uc *userCtl) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, errConv := strconv.ParseInt(idStr, 10, 64)
	if errConv != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, errConv)
		return
	}
	req := &dto.UpdateUserReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	resp, err := uc.userService.UpdateUser(ctx, uint(id), req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserIDNotFound):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserNotFound, err)
			return
		case errors.Is(err, dto.ErrUserTokenInvalidOrExpired):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserTokenInvalidOrExpired, err)
			return
		case errors.Is(err, dto.ErrPermissionDenied):
			http_response.ReturnErrMessage(ctx, http_response.ErrPermissionDenied, err)
			return
		case errors.Is(err, dto.ErrUserAlreadyExists):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserAlreadyExists, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}
	}
	http_response.ReturnSuccessMessage(ctx, resp)
}

func (uc *userCtl) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, errConv := strconv.ParseInt(idStr, 10, 64)
	if errConv != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, errConv)
		return
	}
	resp, err := uc.userService.DeleteUser(ctx, uint(id))
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserIDNotFound):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserNotFound, err)
			return
		case errors.Is(err, dto.ErrUserTokenInvalidOrExpired):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserTokenInvalidOrExpired, err)
			return
		case errors.Is(err, dto.ErrPermissionDenied):
			http_response.ReturnErrMessage(ctx, http_response.ErrPermissionDenied, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}
	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}

}

func (uc *userCtl) List(ctx *gin.Context) {
	req := &dto.ListUserReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	resp, err := uc.userService.List(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserTokenInvalidOrExpired):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserTokenInvalidOrExpired, err)
			return
		case errors.Is(err, dto.ErrPermissionDenied):
			http_response.ReturnErrMessage(ctx, http_response.ErrPermissionDenied, err)
			return
		case errors.Is(err, dto.ErrStrconv):
			http_response.ReturnErrMessage(ctx, http_response.ErrStrconv, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}
	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}

}
