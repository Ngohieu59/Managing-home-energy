package controller

import (
	http_response "Managing-home-energy/cmd/api/controller/http-response"
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type EBillsController interface {
	EAmount(ctx *gin.Context)
	ReportMonthly(ctx *gin.Context)
	EstimateEBill(ctx *gin.Context)
}

type EBillsCtl struct {
	EBillsService service.EBillsService
}

func NewEBillsController(di *do.Injector) EBillsController {
	return &EBillsCtl{
		EBillsService: do.MustInvoke[service.EBillsService](di),
	}
}

func (ec *EBillsCtl) EAmount(ctx *gin.Context) {
	req := &dto.EBillMoneyReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	resp, err := ec.EBillsService.GetEAmount(ctx, req)

	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserTokenInvalidOrExpired):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserTokenInvalidOrExpired, err)
			return
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
func (ec *EBillsCtl) ReportMonthly(ctx *gin.Context) {
	monthStr := ctx.Query("month")
	if monthStr == "" {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("month is empty"))
		return
	}
	resp, err := ec.EBillsService.ReportMonthlyUsageComparison(ctx, monthStr)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrUserTokenInvalidOrExpired):
			http_response.ReturnErrMessage(ctx, http_response.ErrUserTokenInvalidOrExpired, err)
			return
		case errors.Is(err, dto.ErrStrconv):
			http_response.ReturnErrMessage(ctx, http_response.ErrStrconv, err)
			return
		case errors.Is(err, dto.ErrDataFormatWrong):
			http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("invalid Month"))
		default:
			http_response.ReturnErrMessage(ctx, http_response.ErrInternalServerError, err)
			return
		}
	} else {
		http_response.ReturnSuccessMessage(ctx, resp)
	}
}

func (ec *EBillsCtl) EstimateEBill(ctx *gin.Context) {
	req := &dto.EBillMoneyReq{}
	errReq := ctx.ShouldBind(req)
	if errReq != nil {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errReq)
		return
	}
	if req.Electric < 0 {
		http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("electric value must be greater than 0"))
		return
	}
	resp, err := ec.EBillsService.EstimateEBill(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, dto.ErrDataFormatWrong):
			http_response.ReturnErrMessage(ctx, http_response.ErrDataFormatWrong, errors.New("start date format error, format: dd-mm-yyyy"))
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
