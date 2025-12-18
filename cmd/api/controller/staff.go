package controller

import (
	http_response "Managing-home-energy/cmd/api/controller/http-response"
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type StaffController interface {
	GetReportList(ctx *gin.Context)
	GetReportEUsed(ctx *gin.Context)
}

type StaffCtl struct {
	StaffService service.StaffService
}

func NewStaffController(di *do.Injector) StaffController {
	return &StaffCtl{
		StaffService: do.MustInvoke[service.StaffService](di),
	}
}

func (s *StaffCtl) GetReportList(ctx *gin.Context) {
	req := &dto.StaffReportListReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
	}
	resp, err := s.StaffService.GetReportList(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrDataFormatWrong):
			http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("date format error, format: dd-mm-yyyy"))
			return
		case errors.Is(err, dto.ErrStartAfterEnd):
			http_response.ReturnErrMessage(ctx, http_response.ErrStartAfterEnd, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}

	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}
}

func (s *StaffCtl) GetReportEUsed(ctx *gin.Context) {
	req := &dto.ReportEUseReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
	}
	resp, err := s.StaffService.GetReportEUse(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrDataFormatWrong):
			http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("date format error, format: dd-mm-yyyy"))
			return
		case errors.Is(err, dto.ErrStartAfterEnd):
			http_response.ReturnErrMessage(ctx, http_response.ErrStartAfterEnd, err)
			return
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}
	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}

}
