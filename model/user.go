package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:char(36);unique;column:uuid" json:"uuid"`
	Name       string    `gorm:"column:name"`
	Username   string    `gorm:"type:varchar(100);unique;index:idx_username" json:"username"`
	Age        int       `gorm:"column:age"`
	Pass       string    `gorm:"column:pass"`
	Salt       string    `gorm:"column:salt"`
	Permission string    `gorm:"column:permission;default:user" json:"permission"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == uuid.Nil {
		u.UUID, _ = uuid.NewUUID()
		fmt.Println("Create UUID: ", u.UUID)
	} else {
		fmt.Println("Gone Create UUID, exits UUID: ", u.UUID)
	}
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Pass == "" {
		tx.Model(u).Update("pass", fmt.Sprintf("random-pass-%v", u.ID))
	}
	return
}
