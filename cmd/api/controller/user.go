package controller

import (
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"fmt"
	"net/http"
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
	_ = ctx.ShouldBind(req)
	resp, err := uc.userService.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uc *userCtl) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	req := &dto.UpdateUserReq{}
	_ = ctx.ShouldBind(req)
	resp, err := uc.userService.UpdateUser(ctx, uint(id), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uc *userCtl) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	resp, err := uc.userService.DeleteUser(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Đã xóa thành công %v", resp),
		})
	}

}

func (uc *userCtl) List(ctx *gin.Context) {
	req := &dto.ListUserReq{}
	_ = ctx.ShouldBind(req)
	resp, err := uc.userService.List(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, resp)
	}

}
