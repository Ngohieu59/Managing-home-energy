package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Electricity_used struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"type:varchar(36);unique;column:uuid" json:"uuid"`
	Date_used time.Time `gorm:"column:date_used" json:"date_used"`
	Username  string    `gorm:"type:varchar(100);index:idx_username" json:"username"`
	Elec_used float64   `gorm:"column:elec_used" json:"elec_used"`
}

func (*Electricity_used) TableName() string {
	return "electricity_useds"
}

func (u *Electricity_used) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == uuid.Nil {
		u.UUID, _ = uuid.NewUUID()
		fmt.Println("Create UUID: ", u.UUID)
	} else {
		fmt.Println("Gone Create UUID, exits UUID: ", u.UUID)
	}
	return nil
}
