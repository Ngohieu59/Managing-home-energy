package mysql

import (
	"Managing-home-energy/model"
	"context"
	"time"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type EbillRepository interface {
	FindAllRecordByName(ctx context.Context, name string, startDate time.Time, endDate time.Time) ([]*model.Electricity_used, error)
	FindAllRecordByMonth(ctx context.Context, name string, Month int, Year int) ([]*model.Electricity_used, error)
}

type ebillRepo struct {
	db *gorm.DB
}

func newEbillRepo(di *do.Injector) (EbillRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &ebillRepo{db: db}, nil
}

func (e *ebillRepo) FindAllRecordByName(ctx context.Context, name string, startDate time.Time, endDate time.Time) ([]*model.Electricity_used, error) {
	var eBillsList []*model.Electricity_used
	//fmt.Printf("StartDate: %v, EndDate: %v\n", startDate, endDate)
	err := e.db.WithContext(ctx).
		Where("username = ? AND date_used BETWEEN ? AND ?", name, startDate, endDate).
		Order("date_used ASC").
		Find(&eBillsList).Error
	if err != nil {
		return nil, err
	}
	return eBillsList, nil
}

func (e *ebillRepo) FindAllRecordByMonth(ctx context.Context, name string, Month int, Year int) ([]*model.Electricity_used, error) {
	var eBillsList []*model.Electricity_used

	err := e.db.WithContext(ctx).
		Where("username = ? AND MONTH(date_used) = ? AND YEAR(date_used) = ?", name, Month, Year).
		Order("date_used ASC").
		Find(&eBillsList).Error
	if err != nil {
		return nil, err
	}
	return eBillsList, nil

}
