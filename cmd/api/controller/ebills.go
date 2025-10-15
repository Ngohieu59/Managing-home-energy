package controller

import (
	"Managing-home-energy/dto"
	"Managing-home-energy/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type EBillsController interface {
	EMoney(ctx *gin.Context)
	ReportMonthly(ctx *gin.Context)
}

type ebillsCtl struct {
	eBillsService service.EbillsService
}

func NewEBillsController(di *do.Injector) EBillsController {
	return &ebillsCtl{
		eBillsService: do.MustInvoke[service.EbillsService](di),
	}
}

func (ec *ebillsCtl) EMoney(ctx *gin.Context) {
	req := &dto.EBillMoneyReq{}
	_ = ctx.ShouldBind(req)
	resp, err := ec.eBillsService.GeteBillMoney(ctx, req)
	//fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
func (ec *ebillsCtl) ReportMonthly(ctx *gin.Context) {
	monthStr := ctx.Query("month")

	resp, err := ec.eBillsService.ReportMonthlyUsageComparison(ctx, monthStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
