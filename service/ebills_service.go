package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/repository/mysql"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

const (
	layout = "02-01-2006" // định dạng dd-mm-yyyy
	day    = 31
)

var (
	UnitQuantityL1 = constants.UnitLevel1.Quantity / day
	UnitQuantityL2 = constants.UnitLevel2.Quantity / day
	UnitQuantityL3 = constants.UnitLevel3.Quantity / day
	UnitQuantityL4 = constants.UnitLevel4.Quantity / day
	UnitQuantityL5 = constants.UnitLevel5.Quantity / day
	UnitPriceL1    = constants.UnitLevel1.UnitPrice
	UnitPriceL2    = constants.UnitLevel2.UnitPrice
	UnitPriceL3    = constants.UnitLevel3.UnitPrice
	UnitPriceL4    = constants.UnitLevel4.UnitPrice
	UnitPriceL5    = constants.UnitLevel5.UnitPrice
	UnitPriceL6    = constants.UnitLevel6.UnitPrice
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
	userName, exists := ctx.Get(constants.ClaimUsername)
	if !exists {
		return nil, dto.ErrUserTokenInvalidOrExpired
	}

	StartDate, errS := time.Parse(layout, req.StartDate)
	if errS != nil {
		return nil, dto.ErrDataFormatWrong
	}
	EndDate, errD := time.Parse(layout, req.EndDate)
	if errD != nil {
		return nil, dto.ErrDataFormatWrong
	}

	if StartDate.After(EndDate) {
		return nil, dto.ErrStartAfterEnd
	}

	TotalEUsed, err := e.EBillRepo.FindAllRecordByName(ctx, userName.(string), StartDate, EndDate)
	if err != nil {
		return nil, err
	}

	resp := &dto.EBillMoneyResp{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		ElectUsed: TotalEUsed,
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
	var TotalMoneyBeforeTax = 0.0
	EUsed := req.Electric
	StartDate, errS := time.Parse(layout, req.StartDate)
	if errS != nil {
		return nil, dto.ErrDataFormatWrong
	}
	EndDate, errD := time.Parse(layout, req.EndDate)
	if errD != nil {
		return nil, dto.ErrDataFormatWrong
	}
	if StartDate.After(EndDate) {
		return nil, dto.ErrStartAfterEnd
	}

	days := EndDate.Sub(StartDate).Hours()/24 + 1
	QuantityL1, QuantityL2, QuantityL3, QuantityL4, QuantityL5 := math.Round(UnitQuantityL1*days), math.Round(UnitQuantityL2*days), math.Round(UnitQuantityL3*days), math.Round(UnitQuantityL4*days), math.Round(UnitQuantityL5*days)

	if EUsed <= QuantityL1 {
		TotalMoneyBeforeTax += UnitPriceL1 * EUsed
	} else if EUsed <= (QuantityL2 + QuantityL1) {
		TotalMoneyBeforeTax += QuantityL1*UnitPriceL1 + (EUsed-QuantityL1)*UnitPriceL2
	} else if EUsed <= (QuantityL3 + QuantityL2 + QuantityL1) {
		TotalMoneyBeforeTax += QuantityL1*UnitPriceL1 + QuantityL2*UnitPriceL2 + (EUsed-QuantityL1-QuantityL2)*UnitPriceL3
	} else if EUsed <= (QuantityL1 + QuantityL2 + QuantityL4 + QuantityL3) {
		TotalMoneyBeforeTax = QuantityL1*UnitPriceL1 + QuantityL2*UnitPriceL2 + QuantityL3*UnitPriceL3 + (EUsed-QuantityL1-QuantityL2-QuantityL3)*UnitPriceL4
	} else if EUsed <= (QuantityL1 + QuantityL2 + QuantityL3 + QuantityL4 + QuantityL5) {
		TotalMoneyBeforeTax = QuantityL1*UnitPriceL1 + QuantityL2*UnitPriceL2 + QuantityL3*UnitPriceL3 + QuantityL4*UnitPriceL4 + (EUsed-QuantityL1-QuantityL2-QuantityL3-QuantityL4)*UnitPriceL5
	} else {
		TotalMoneyBeforeTax = QuantityL1*UnitPriceL1 + QuantityL2*UnitPriceL2 + QuantityL3*UnitPriceL3 + QuantityL4*UnitPriceL4 + QuantityL5*UnitPriceL5 + (EUsed-QuantityL1-QuantityL2-QuantityL3-QuantityL4-QuantityL5)*UnitPriceL6
	}
	TotalMoney := math.Round(TotalMoneyBeforeTax * (1 + constants.Taxt/100))

	resp := &dto.EBillMoneyResp{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Money:     TotalMoney,
		ElectUsed: EUsed,
	}
	return resp, nil
}
