package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID         uuid.UUID `gorm:"type:char(36);unique;column:uuid" json:"uuid"`
	Name         string    `gorm:"column:name"`
	Username     string    `gorm:"type:varchar(100);unique;index:idx_username" json:"username"`
	Age          int       `gorm:"column:age"`
	Pass         string    `gorm:"column:pass"`
	Salt         string    `gorm:"column:salt"`
	Permission   string    `gorm:"column:permission;default:user" json:"permission"`
	Type         string    `gorm:"column:type;type:enum('family','business','industrial','administrative');not null"`
	VoltageLevel string    `gorm:"column:voltageLevel;type:enum('low','medium','high');default:null"`
	// 🏠 Address information
	HouseNumber string `gorm:"column:houseNumber;size:50"` // Số nhà
	Ward        string `gorm:"column:ward;size:20"`        // Phường
	City        string `gorm:"column:city;size:20"`        // Thành phố
}

func (*User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(ctx *gorm.DB) (err error) {
	if u.UUID == uuid.Nil {
		u.UUID, _ = uuid.NewUUID()
		fmt.Println("Create UUID: ", u.UUID)
	} else {
		fmt.Println("Gone Create UUID, exits UUID: ", u.UUID)
	}
	return nil
}

func (u *User) AfterCreate(ctx *gorm.DB) (err error) {
	if u.Pass == "" {
		ctx.Model(u).Update("pass", fmt.Sprintf("random-pass-%v", u.ID))
	}
	return
}
