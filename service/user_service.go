package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/model"
	"Managing-home-energy/repository/mysql"
	"Managing-home-energy/utils"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/do"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.User, error)
	UpdateUser(ctx *gin.Context, id uint, req *dto.UpdateUserReq) (*dto.User, error)
	DeleteUser(ctx *gin.Context, id uint) (string, error)
	List(ctx *gin.Context, req *dto.ListUserReq) (*dto.ListUserResponse, error)
}

type userServiceImpl struct {
	userRepo mysql.UserRepository
}

func newUserService(di *do.Injector) (UserService, error) {
	return &userServiceImpl{
		userRepo: do.MustInvoke[mysql.UserRepository](di),
	}, nil
}

func (s *userServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.User, error) {
	existedUser, _ := s.userRepo.FindByName(ctx, req.Username)
	if existedUser != nil && existedUser.Username != "" {
		return nil, errors.New("username already exist")
	}

	if req.Name == "" || req.Username == "" || req.Password == "" {
		return nil, errors.New("please fill all fields")
	}

	if req.Age <= 0 {
		return nil, errors.New("please set age > 0")
	}

	user := &model.User{
		Name:     req.Name,
		Age:      req.Age,
		Pass:     req.Password,
		Username: req.Username,
	}
	salt, _ := utils.GenerateSalt()
	hashedPass := utils.HashPassword(req.Password, salt)
	user.Pass = hashedPass
	user.Salt = salt
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	userRes := s.convertToUserDto(user)
	return userRes, nil
}

func (s *userServiceImpl) convertToUserDto(user *model.User) *dto.User {
	res := &dto.User{
		ID:         user.ID,
		Name:       user.Name,
		Age:        user.Age,
		Username:   user.Username,
		Permission: user.Permission,
	}
	return res
}

func (s *userServiceImpl) UpdateUser(ctx *gin.Context, id uint, req *dto.UpdateUserReq) (*dto.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	userID, existsID := ctx.Get(constants.ClaimUserId)
	userP, existsP := ctx.Get(constants.ClaimPermission)

	if !existsID || !existsP {
		return nil, errors.New("claim user not found")
	}
	if userP == "user" && userID != user.ID {
		return nil, errors.New("You don't have the permission")
	}
	if req.Username != "" {
		existedUser, _ := s.userRepo.FindByName(ctx, req.Username)
		if existedUser != nil && existedUser.Username != "" {
			return nil, errors.New("username already exist")
		} else {
			user.Username = req.Username
		}
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Password != "" {
		salt, _ := utils.GenerateSalt()
		hashedPass := utils.HashPassword(req.Password, salt)
		user.Pass = hashedPass
		user.Salt = salt
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	if user.Permission == "admin" && req.Permission != "" {
		user.Permission = req.Permission
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	userRes := s.convertToUserDto(user)
	return userRes, nil
}

func (s *userServiceImpl) DeleteUser(ctx *gin.Context, id uint) (string, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return "", err
	}
	userID, existsID := ctx.Get(constants.ClaimUserId)
	userP, existsP := ctx.Get(constants.ClaimPermission)
	if !existsID || !existsP {
		return "", errors.New("claim user not found")
	}
	if userP == "user" && userID != user.ID {
		return "", errors.New("You don't have the permission")
	}

	err = s.userRepo.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (s *userServiceImpl) List(ctx *gin.Context, req *dto.ListUserReq) (*dto.ListUserResponse, error) {
	userP, existsP := ctx.Get(constants.ClaimPermission)
	if !existsP {
		return nil, errors.New("claim user not found")
	}
	if userP != "admin" {
		return nil, errors.New("You don't have the permission")
	}
	limit, _ := strconv.Atoi(req.Limit)
	offset, _ := strconv.Atoi(req.Offset)
	orderBy := req.OrderBy

	users, err := s.userRepo.List(ctx, limit, offset, orderBy)
	if err != nil {
		return nil, err
	}
	return users, nil
}
