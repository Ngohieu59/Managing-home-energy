package mysql

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/model"
	"Managing-home-energy/utils"
	"context"
	"time"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type EBillRepository interface {
	FindAllRecordByName(ctx context.Context, name string, startDate time.Time, endDate time.Time) (float64, error)
	FindAllRecordByMonth(ctx context.Context, name string, Month int, Year int) (float64, error)
	CalcEUsed(ctx context.Context, username string, startDate time.Time, endDate time.Time, userType string) (*dto.ListEUsedResponse, error)
}

type EBillRepo struct {
	db *gorm.DB
}

func newEBillRepo(di *do.Injector) (EBillRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &EBillRepo{db: db}, nil
}

func (e *EBillRepo) FindAllRecordByName(ctx context.Context, name string, startDate time.Time, endDate time.Time) (float64, error) {
	var total = 0.0
	err := e.db.WithContext(ctx).
		Model(&model.Electricity_used{}).
		Select("SUM(elec_used)").
		Where("username = ? AND date_used BETWEEN ? AND ?", name, startDate, endDate).
		Scan(&total).Error
	if err != nil {
		return total, err
	}
	return total, nil
}

func (e *EBillRepo) CalcEUsed(ctx context.Context, username string, startDate time.Time, endDate time.Time, userType string) (*dto.ListEUsedResponse, error) {
	var total = 0.0
	var totalEUsedLow = 0.0
	var totalEUsedNormal = 0.0
	var totalEUsedHigh = 0.0
	baseQuery := e.db.WithContext(ctx).
		Model(&model.Electricity_used{}).
		Select("COALESCE(SUM(elec_used), 0)").
		Where("username = ? AND date_used BETWEEN ? AND ?", username, startDate, endDate)
	err := baseQuery.Scan(&total).Error
	if err != nil {
		return nil, err
	}
	if userType == constants.TypeBusiness || userType == constants.TypeIndustrial {
		orSQLNormal, argsNormal := utils.StringSQL(constants.TariffHours.Normal)
		normalQuery := e.db.WithContext(ctx).
			Model(&model.Electricity_used{}).
			Select("COALESCE(SUM(elec_used), 0)").
			Where("username = ? AND date_used BETWEEN ? AND ? AND deleted_at IS NULL", username, startDate, endDate).
			Where(orSQLNormal, argsNormal...)
		errN := normalQuery.Scan(&totalEUsedNormal).Error
		if errN != nil {
			return nil, errN
		}

		orSQLLow, argsLow := utils.StringSQL(constants.TariffHours.Low)
		lowQuery := e.db.WithContext(ctx).
			Model(&model.Electricity_used{}).
			Select("COALESCE(SUM(elec_used), 0)").
			Where("username = ? AND date_used BETWEEN ? AND ? AND deleted_at IS NULL", username, startDate, endDate).
			Where(orSQLLow, argsLow...)
		errL := lowQuery.Scan(&totalEUsedLow).Error
		if errL != nil {
			return nil, errL
		}

		orSQLHigh, argsHigh := utils.StringSQL(constants.TariffHours.High)
		HighQuery := e.db.WithContext(ctx).
			Model(&model.Electricity_used{}).
			Select("COALESCE(SUM(elec_used), 0)").
			Where("username = ? AND date_used BETWEEN ? AND ? AND deleted_at IS NULL", username, startDate, endDate).
			Where(orSQLHigh, argsHigh...)
		errH := HighQuery.Scan(&totalEUsedHigh).Error
		if errH != nil {
			return nil, errH
		}
	}
	ListResp := &dto.ListEUsedResponse{
		Total:  total,
		Normal: totalEUsedNormal,
		Low:    totalEUsedLow,
		High:   totalEUsedHigh,
	}
	return ListResp, nil
}

func (e *EBillRepo) FindAllRecordByMonth(ctx context.Context, name string, Month int, Year int) (float64, error) {
	var total = 0.0

	err := e.db.WithContext(ctx).
		Model(&model.Electricity_used{}).
		Select("SUM(elec_used)").
		Where("username = ? AND MONTH(date_used) = ? AND YEAR(date_used) = ?", name, Month, Year).
		Order("date_used ASC").
		Scan(&total).Error
	if err != nil {
		return total, err
	}
	return total, nil

}
