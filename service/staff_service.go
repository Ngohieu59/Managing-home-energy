package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/repository/mysql"
	"Managing-home-energy/utils"
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

const day = 31

var (
	NumFamilyQuantityL1 = 0
	NumFamilyQuantityL2 = 0
	NumFamilyQuantityL3 = 0
	NumFamilyQuantityL4 = 0
	NumFamilyQuantityL5 = 0
	NumFamilyQuantityL6 = 0

	SumBusinessLevel1 = 0.0
	SumBusinessLevel2 = 0.0
	SumBusinessLevel3 = 0.0

	SumIndustrialLevel1 = 0.0
	SumIndustrialLevel2 = 0.0
	SumIndustrialLevel3 = 0.0
	SumIndustrialLevel4 = 0.0

	SumAdministrativeLevel1 = 0.0
	SumAdministrativeLevel2 = 0.0

	BusinessUnitVolLevel1 = constants.UnitBusinessLowLevel1.Quantity
	BusinessUnitVolLevel2 = constants.UnitBusinessHighLevel1.Quantity

	IndustrialUnitVolLevel1 = constants.UnitIndustrialLowLevel1.Quantity
	IndustrialUnitVolLevel2 = constants.UnitIndustrialLowLevel2.Quantity
	IndustrialUnitVolLevel3 = constants.UnitIndustrialLowLevel2.Quantity

	AdministrativeUnitVolLevel1 = constants.UnitAdministrativeLowLevel1.Quantity
)

type StaffService interface {
	GetReportList(ctx *gin.Context, req *dto.StaffReportListReq) (*dto.StaffReportListResp, error)
	GetReportEUse(ctx *gin.Context, req *dto.ReportEUseReq) (*dto.ReportEUseResp, error)
}

type StaffServiceImpl struct {
	userRepo  mysql.UserRepository
	EBillRepo mysql.EBillRepository
}

func newStaffService(di *do.Injector) (StaffService, error) {
	return &StaffServiceImpl{
		userRepo:  do.MustInvoke[mysql.UserRepository](di),
		EBillRepo: do.MustInvoke[mysql.EBillRepository](di),
	}, nil
}

func (s *StaffServiceImpl) GetReportList(ctx *gin.Context, req *dto.StaffReportListReq) (*dto.StaffReportListResp, error) {
	StartDate, EndDate, errD := utils.CheckDate(req.StartDate, req.EndDate)
	if errD != nil {
		return nil, errD
	}
	ListUser, err := s.userRepo.ListUser(ctx, req.Type, req.City, req.Ward)
	if err != nil {
		return nil, err
	}
	if len(ListUser.Data) == 0 {
		return &dto.StaffReportListResp{
			Filter:   req,
			ListUser: make([]*dto.UserReport, 0),
		}, nil
	}

	days := EndDate.Sub(StartDate).Hours()/24 + 1
	var userReport []*dto.UserReport
	for _, user := range ListUser.Data {
		TotalMoneyBeforeTax := 0.0
		TotalEUsed, err := s.EBillRepo.CalcEUsed(ctx, user.Username, StartDate, EndDate, user.Type)
		if err != nil {
			return nil, err
		}
		if TotalEUsed.Total >= req.Threshold {
			switch user.Type {
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
			userReport = append(userReport, &dto.UserReport{
				ID:        user.ID,
				Name:      user.Name,
				Username:  user.Username,
				Money:     TotalMoney,
				ElectUsed: TotalEUsed.Total,
			})
		}
	}
	resp := &dto.StaffReportListResp{
		Filter:   req,
		ListUser: userReport,
	}
	return resp, nil
}

func (s *StaffServiceImpl) GetReportEUse(ctx *gin.Context, req *dto.ReportEUseReq) (*dto.ReportEUseResp, error) {
	StartDate, EndDate, errD := utils.CheckDate(req.StartDate, req.EndDate)
	if errD != nil {
		return nil, errD
	}

	ListUser, err := s.userRepo.ListUser(ctx, req.Type, req.City, req.Ward)
	if err != nil {
		return nil, err
	}
	days := EndDate.Sub(StartDate).Hours()/24 + 1

	QuantityL1, QuantityL2, QuantityL3, QuantityL4, QuantityL5 := math.Round(utils.UnitFamilyQuantityL1*days), math.Round(utils.UnitFamilyQuantityL2*days), math.Round(utils.UnitFamilyQuantityL3*days), math.Round(utils.UnitFamilyQuantityL4*days), math.Round(utils.UnitFamilyQuantityL5*days)

	BusinessVolLevel1 := BusinessUnitVolLevel1 / day * days
	BusinessVolLevel2 := BusinessUnitVolLevel2 / day * days

	IndustrialVolLevel1 := IndustrialUnitVolLevel1 / day * days
	IndustrialVolLevel2 := IndustrialUnitVolLevel2 / day * days
	IndustrialVolLevel3 := IndustrialUnitVolLevel3 / day * days

	AdministrativeVolLevel1 := AdministrativeUnitVolLevel1 / day * days
	for _, user := range ListUser.Data {
		if user.Type == req.Type {
			TotalEUsed, err := s.EBillRepo.CalcEUsed(ctx, user.Username, StartDate, EndDate, user.Type)
			if err != nil {
				return nil, err
			}
			switch req.Type {

			case constants.TypeFamily:
				{
					if TotalEUsed.Total <= QuantityL1 {
						NumFamilyQuantityL1 += 1
					} else if TotalEUsed.Total <= QuantityL2 {
						NumFamilyQuantityL2 += 1
					} else if TotalEUsed.Total <= QuantityL3 {
						NumFamilyQuantityL3 += 1
					} else if TotalEUsed.Total <= QuantityL4 {
						NumFamilyQuantityL4 += 1
					} else if TotalEUsed.Total <= QuantityL5 {
						NumFamilyQuantityL5 += 1
					} else {
						NumFamilyQuantityL6 += 1
					}
				}

			case constants.TypeBusiness:
				{
					if TotalEUsed.Total <= BusinessVolLevel1 {
						SumBusinessLevel1 += TotalEUsed.High
					} else if TotalEUsed.Total <= BusinessVolLevel2 {
						SumBusinessLevel2 += TotalEUsed.High
					} else {
						SumBusinessLevel3 += TotalEUsed.High
					}
				}

			case constants.TypeIndustrial:
				if TotalEUsed.Total <= IndustrialVolLevel1 {
					SumIndustrialLevel1 += TotalEUsed.High
				} else if TotalEUsed.Total <= IndustrialVolLevel2 {
					SumIndustrialLevel2 += TotalEUsed.High
				} else if TotalEUsed.Total <= IndustrialVolLevel3/day*days {
					SumIndustrialLevel3 += TotalEUsed.High
				} else {
					SumIndustrialLevel4 += TotalEUsed.High
				}

			case constants.TypeAdministrative:
				if TotalEUsed.Total <= AdministrativeVolLevel1 {
					SumAdministrativeLevel1 += TotalEUsed.Total
				} else {
					SumAdministrativeLevel2 += TotalEUsed.Total
				}
			}
		}
	}
	resp := &dto.ReportEUseResp{
		Filter: req,
	}
	switch req.Type {
	case constants.TypeFamily:
		resp.Data = &dto.FamilyResp{
			TotalMember: []*dto.TotalLevel{
				{
					VolLevel: "Level1",
					NumUsers: NumFamilyQuantityL1,
				},
				{
					VolLevel: "Level2",
					NumUsers: NumFamilyQuantityL2,
				},
				{
					VolLevel: "Level3",
					NumUsers: NumFamilyQuantityL3,
				},
				{
					VolLevel: "Level4",
					NumUsers: NumFamilyQuantityL4,
				},
				{
					VolLevel: "Level5",
					NumUsers: NumFamilyQuantityL5,
				},
				{
					VolLevel: "Level6",
					NumUsers: NumFamilyQuantityL6,
				},
			},
		}

	case constants.TypeBusiness:
		resp.Data = &dto.VolResp{
			Level: []*dto.Vol{
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp dưới %v", BusinessUnitVolLevel1),
					ElectUsed: SumBusinessLevel1,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp từ %v đến dưới %v", BusinessUnitVolLevel1, BusinessUnitVolLevel2),
					ElectUsed: SumBusinessLevel2,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp trên %v", BusinessUnitVolLevel2),
					ElectUsed: SumBusinessLevel3,
				},
			},
		}
	case constants.TypeIndustrial:

		resp.Data = &dto.VolResp{
			Level: []*dto.Vol{
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp dưới %v", IndustrialUnitVolLevel1),
					ElectUsed: SumIndustrialLevel1,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp từ %v đến dưới %v", IndustrialUnitVolLevel1, IndustrialUnitVolLevel2),
					ElectUsed: SumIndustrialLevel2,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp từ %v đến dưới %v", IndustrialUnitVolLevel2, IndustrialUnitVolLevel3),
					ElectUsed: SumIndustrialLevel3,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp trên %v", IndustrialUnitVolLevel3),
					ElectUsed: SumIndustrialLevel4,
				},
			},
		}

	case constants.TypeAdministrative:

		resp.Data = &dto.VolResp{
			Level: []*dto.Vol{
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp dưới %v", AdministrativeUnitVolLevel1),
					ElectUsed: SumAdministrativeLevel1,
				},
				{
					VolLevel:  fmt.Sprintf("Cấp điện áp trên %v", AdministrativeUnitVolLevel1),
					ElectUsed: SumAdministrativeLevel2,
				},
			},
		}

	}

	return resp, nil
}
