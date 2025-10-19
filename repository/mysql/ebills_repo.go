package mysql

import (
	"Managing-home-energy/model"
	"context"
	"time"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type EBillRepository interface {
	FindAllRecordByName(ctx context.Context, name string, startDate time.Time, endDate time.Time) (float64, error)
	FindAllRecordByMonth(ctx context.Context, name string, Month int, Year int) (float64, error)
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
	//fmt.Printf("StartDate: %v, EndDate: %v\n", startDate, endDate)
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
