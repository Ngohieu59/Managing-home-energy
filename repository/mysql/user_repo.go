package mysql

import (
	"Managing-home-energy/dto"
	"Managing-home-energy/model"
	"context"
	"errors"

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
	ListUser(ctx context.Context, userType string, city string, ward string) (*dto.ListUserResponse, error)
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
		return nil, dto.ErrUserIDNotFound
	}
	return &user, nil
}

func (u *userRepo) FindByName(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	err := u.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("not found user with that username")
	}

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
	if OrderBy == "" {
		OrderBy = "created_at desc"
	}
	query := u.db.WithContext(ctx)
	if limit > 0 {
		query = query.Limit(limit).Offset(offset).Order(OrderBy)
	} else {
		query = query.Offset(offset).Order(OrderBy)
	}

	err := query.Find(&userList).Error

	if err != nil {
		return nil, err
	}
	resp := &dto.ListUserResponse{
		Data: userList,
	}

	return resp, nil
}

func (u *userRepo) ListUser(ctx context.Context, userType string, city string, ward string) (*dto.ListUserResponse, error) {
	var userList []*dto.User
	query := u.db.WithContext(ctx).Model(&dto.User{})

	if userType != "" {
		query = query.Where("type = ?", userType)
	}

	if city != "" {
		query = query.Where("city = ?", city)
	}

	if ward != "" {
		query = query.Where("ward = ?", ward)
	}

	if err := query.Find(&userList).Error; err != nil {
		return nil, err
	}
	resp := &dto.ListUserResponse{
		Data: userList,
	}

	return resp, nil
}
