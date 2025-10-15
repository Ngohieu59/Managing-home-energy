package mysql

import (
	"Managing-home-energy/dto"
	"Managing-home-energy/model"
	"context"
	"fmt"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
	FindByName(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit int, offset int, OrderBy string) (*dto.ListUserResponse, error)
}
type userRepo struct {
	db *gorm.DB
}

func newUserRepo(di *do.Injector) (UserRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &userRepo{db: db}, nil
}

func (u *userRepo) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindByName(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	err := u.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Not found user with name %v", username)
	}
	//fmt.Println(user)
	return user, nil
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Model(user).Updates(user).Error
}

func (u *userRepo) Delete(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).Unscoped().Delete(&model.User{}, id).Error
}

func (u *userRepo) List(ctx context.Context, limit int, offset int, OrderBy string) (*dto.ListUserResponse, error) {
	var userList []*dto.User
	err := u.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order(OrderBy).
		Find(&userList).Error
	if err != nil {
		return nil, err
	}
	resp := &dto.ListUserResponse{
		Data: userList,
	}

	fmt.Printf("data all user %v", resp)
	return resp, nil
}
