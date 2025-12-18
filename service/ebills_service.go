package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/repository/mysql"
	"Managing-home-energy/utils"
	"math"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type EBillsService interface {
	GetEAmount(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EBillMoneyResp, error)
	ReportMonthlyUsageComparison(ctx *gin.Context, Month string) (*dto.ReportMonthlyResp, error)
	EstimateEBill(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EBillMoneyResp, error)
}

type EBillServiceImpl struct {
	EBillRepo mysql.EBillRepository
}

func newEBillService(di *do.Injector) (EBillsService, error) {
	return &EBillServiceImpl{
		EBillRepo: do.MustInvoke[mysql.EBillRepository](di),
	}, nil
}

func (e *EBillServiceImpl) GetEAmount(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EBillMoneyResp, error) {
	var TotalMoneyBeforeTax = 0.0
	userName, exists := ctx.Get(constants.ClaimUsername)
	if !exists {
		return nil, dto.ErrUserTokenInvalidOrExpired
	}
	userType, exist := ctx.Get(constants.ClaimUserType)
	if !exist {
		return nil, dto.ErrUserTokenInvalidOrExpired
	}
	StartDate, EndDate, errD := utils.CheckDate(req.StartDate, req.EndDate)
	if errD != nil {
		return nil, errD
	}

	TotalEUsed, err := e.EBillRepo.CalcEUsed(ctx, userName.(string), StartDate, EndDate, userType.(string))
	if err != nil {
		return nil, err
	}
	days := EndDate.Sub(StartDate).Hours()/24 + 1
	switch userType.(string) {
	case constants.TypeFamily:
		TotalMoneyBeforeTax = utils.MoneyFamily(TotalEUsed.Total, days)
	case constants.TypeBusiness:

		TotalMoneyBeforeTax = utils.MoneyBusiness(TotalEUsed, days)
	case constants.TypeIndustrial:
		TotalMoneyBeforeTax = utils.MoneyIndustrial(TotalEUsed, days)
	case constants.TypeAdministrative:
		TotalMoneyBeforeTax = utils.MoneyAdministrative(TotalEUsed.Total, days)
	}
	TotalMoney := math.Round(TotalMoneyBeforeTax * (1 + constants.Taxt/100))
	resp := &dto.EBillMoneyResp{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		ElectUsed: TotalEUsed.Total,
		Money:     TotalMoney,
	}
	return resp, nil
}

func (e *EBillServiceImpl) ReportMonthlyUsageComparison(ctx *gin.Context, Month string) (*dto.ReportMonthlyResp, error) {
	userName, exists := ctx.Get(constants.ClaimUsername)
	if !exists {
		return nil, dto.ErrUserTokenInvalidOrExpired
	}
	MonthInt, Err := strconv.Atoi(Month)
	if Err != nil {
		return nil, dto.ErrStrconv
	}
	if MonthInt < 1 || MonthInt > 12 {
		return nil, dto.ErrDataFormatWrong
	}

	currentYear := time.Now().Year()

	TotalElectUsedCurrentYear, errC := e.EBillRepo.FindAllRecordByMonth(ctx, userName.(string), MonthInt, currentYear)
	if errC != nil {
		return nil, errC
	}
	TotalElectUsedLastYear, errL := e.EBillRepo.FindAllRecordByMonth(ctx, userName.(string), MonthInt, currentYear-1)
	if errL != nil {
		return nil, errL
	}

	reps := &dto.ReportMonthlyResp{
		ThisYear: &dto.ReportMonthly{
			Month:     MonthInt,
			Year:      currentYear,
			ElectUsed: TotalElectUsedCurrentYear,
		},
		LastYear: &dto.ReportMonthly{
			Month:     MonthInt,
			Year:      currentYear - 1,
			ElectUsed: TotalElectUsedLastYear,
		},
	}
	return reps, nil
}

func (e *EBillServiceImpl) EstimateEBill(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EBillMoneyResp, error) {
	userType, exists := ctx.Get(constants.ClaimUsername)
	if !exists {
		return nil, dto.ErrUserTokenInvalidOrExpired
	}
	if userType != constants.TypeFamily {
		return nil, dto.ErrPermissionDenied
	}
	var TotalMoney = 0.0
	EUsed := req.Electric
	StartDate, EndDate, errD := utils.CheckDate(req.StartDate, req.EndDate)
	if errD != nil {
		return nil, errD
	}
	days := EndDate.Sub(StartDate).Hours()/24 + 1

	TotalMoney = utils.MoneyFamily(EUsed, days)
	resp := &dto.EBillMoneyResp{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Money:     TotalMoney,
		ElectUsed: EUsed,
	}
	return resp, nil
}
