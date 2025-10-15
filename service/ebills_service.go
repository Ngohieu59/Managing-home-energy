package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/repository/mysql"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

const (
	layout = "02-01-2006" // định dạng dd-mm-yyyy
)

type EbillsService interface {
	GeteBillMoney(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EbillMoneyResp, error)
	ReportMonthlyUsageComparison(ctx *gin.Context, Month string) (*dto.ReportMonthlyResp, error)
}

type EbillServiceImpl struct {
	ebillRepo mysql.EbillRepository
}

func newEbillService(di *do.Injector) (EbillsService, error) {
	return &EbillServiceImpl{
		ebillRepo: do.MustInvoke[mysql.EbillRepository](di),
	}, nil
}

func (e *EbillServiceImpl) GeteBillMoney(ctx *gin.Context, req *dto.EBillMoneyReq) (*dto.EbillMoneyResp, error) {
	var TotalElecUsed = 0.0
	var TotalMoneyBeforeTax = 0.0
	userName, _ := ctx.Get(constants.ClaimUsername)
	StartDate, errS := time.Parse(layout, req.StartDate)
	//fmt.Printf("StartDate %s\n", StartDate)
	if errS != nil {
		return nil, errors.New("start date format error, format: dd-mm-yyyy")
	}
	EndDate, errD := time.Parse(layout, req.EndDate)
	//fmt.Printf("EndDate %s\n", EndDate)
	if errD != nil {
		return nil, errors.New("end date format error, format: dd-mm-yyyy")
	}
	if StartDate.After(EndDate) {
		return nil, errors.New("start date after end date")
	}
	//fmt.Printf("Name %v", userName)
	//fmt.Printf("StartDate: %s, EndDate: %s\n", StartDate, EndDate)
	AllERecord, err := e.ebillRepo.FindAllRecordByName(ctx, userName.(string), StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(AllERecord); i++ {
		TotalElecUsed += AllERecord[i].Elec_used
	}
	//fmt.Printf("Luong dien su dung: %v", TotalElecUsed)
	if TotalElecUsed <= constants.UnitLevel1.Quantity {
		TotalMoneyBeforeTax += constants.UnitLevel1.UnitPrice * TotalElecUsed
	} else if TotalElecUsed <= (constants.UnitLevel2.Quantity + constants.UnitLevel1.Quantity) {
		TotalMoneyBeforeTax += constants.UnitLevel1.Quantity*constants.UnitLevel1.UnitPrice + (TotalElecUsed-constants.UnitLevel1.Quantity)*constants.UnitLevel2.UnitPrice
	} else if TotalElecUsed <= (constants.UnitLevel3.Quantity + constants.UnitLevel2.Quantity + constants.UnitLevel1.Quantity) {
		TotalMoneyBeforeTax += constants.UnitLevel1.Quantity*constants.UnitLevel1.UnitPrice + constants.UnitLevel2.Quantity*constants.UnitLevel2.UnitPrice + (TotalElecUsed-constants.UnitLevel1.Quantity-constants.UnitLevel2.Quantity)*constants.UnitLevel3.UnitPrice
	} else if TotalElecUsed <= (constants.UnitLevel1.Quantity + constants.UnitLevel2.Quantity + constants.UnitLevel4.Quantity + constants.UnitLevel3.Quantity) {
		TotalMoneyBeforeTax = constants.UnitLevel1.Quantity*constants.UnitLevel1.UnitPrice + constants.UnitLevel2.Quantity*constants.UnitLevel2.UnitPrice + constants.UnitLevel3.Quantity*constants.UnitLevel3.UnitPrice + (TotalElecUsed-constants.UnitLevel1.Quantity-constants.UnitLevel2.Quantity-constants.UnitLevel3.Quantity)*constants.UnitLevel4.UnitPrice
	} else if TotalElecUsed <= (constants.UnitLevel1.Quantity + constants.UnitLevel2.Quantity + constants.UnitLevel3.Quantity + constants.UnitLevel4.Quantity + constants.UnitLevel5.Quantity) {
		TotalMoneyBeforeTax = constants.UnitLevel1.Quantity*constants.UnitLevel1.UnitPrice + constants.UnitLevel2.Quantity*constants.UnitLevel2.UnitPrice + constants.UnitLevel3.Quantity*constants.UnitLevel3.UnitPrice + constants.UnitLevel4.Quantity*constants.UnitLevel4.UnitPrice + (TotalElecUsed-constants.UnitLevel1.Quantity-constants.UnitLevel2.Quantity-constants.UnitLevel3.Quantity-constants.UnitLevel4.Quantity)*constants.UnitLevel5.UnitPrice
	} else {
		TotalMoneyBeforeTax = constants.UnitLevel1.Quantity*constants.UnitLevel1.UnitPrice + constants.UnitLevel2.Quantity*constants.UnitLevel2.UnitPrice + constants.UnitLevel3.Quantity*constants.UnitLevel3.UnitPrice + constants.UnitLevel4.Quantity*constants.UnitLevel4.UnitPrice + constants.UnitLevel5.Quantity*constants.UnitLevel5.UnitPrice + (TotalElecUsed-constants.UnitLevel1.Quantity-constants.UnitLevel2.Quantity-constants.UnitLevel3.Quantity-constants.UnitLevel4.Quantity-constants.UnitLevel5.Quantity)*constants.UnitLevel6.UnitPrice
	}
	TotalMoney := TotalMoneyBeforeTax * (1 + constants.Taxt/100)
	//fmt.Printf("TotalMoney: %v, TotalMoneyBeforeTax: %v", TotalMoney, TotalMoneyBeforeTax)
	resp := &dto.EbillMoneyResp{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Money:     TotalMoney,
		ElectUsed: TotalElecUsed,
	}
	return resp, nil
}

func (e *EbillServiceImpl) ReportMonthlyUsageComparison(ctx *gin.Context, Month string) (*dto.ReportMonthlyResp, error) {
	var TotalElectUsedCurrentYear = 0.0
	var TotalElectUsedLastYear = 0.0
	if Month == "" {
		return nil, errors.New("month is empty")
	}
	userName, _ := ctx.Get(constants.ClaimUsername)
	MonthInt, Err := strconv.Atoi(Month)
	if Err != nil {
		return nil, Err
	}
	if MonthInt < 1 || MonthInt > 12 {
		return nil, errors.New("invalid Month")
	}

	currentYear := time.Now().Year()

	ReportCurrentYear, errC := e.ebillRepo.FindAllRecordByMonth(ctx, userName.(string), MonthInt, currentYear)
	if errC != nil {
		return nil, errC
	}
	ReportLastYear, errL := e.ebillRepo.FindAllRecordByMonth(ctx, userName.(string), MonthInt, currentYear-1)
	if errL != nil {
		return nil, errL
	}
	for i := 0; i < len(ReportCurrentYear); i++ {
		TotalElectUsedCurrentYear += ReportCurrentYear[i].Elec_used
	}
	for i := 0; i < len(ReportLastYear); i++ {
		TotalElectUsedLastYear += ReportLastYear[i].Elec_used
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
