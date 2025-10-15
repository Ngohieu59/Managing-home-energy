package service

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"Managing-home-energy/log"
	"Managing-home-energy/repository/mysql"
	"Managing-home-energy/utils"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

type AuthService interface {
	PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error)
}

type authServiceImpl struct {
	userRepo       mysql.UserRepository
	jwtUtil        utils.JWTUtil
	tenantId       string
	AuthCodeSymKey string
}

func newAuthService(di *do.Injector) (AuthService, error) {
	authServiceImpl := &authServiceImpl{
		userRepo: do.MustInvoke[mysql.UserRepository](di),
		jwtUtil:  do.MustInvoke[utils.JWTUtil](di),
	}
	authServiceImpl.tenantId = constants.TenantID

	key, err := authServiceImpl.authCodeSymKeyGenerate(authServiceImpl.tenantId)
	if err != nil {
		return nil, err
	}
	authServiceImpl.AuthCodeSymKey = key
	return authServiceImpl, err

}

func (s *authServiceImpl) PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	hashedPass := utils.HashPassword(req.Password, user.Salt)
	//fmt.Printf("Password : %s\n", req.Password)
	//fmt.Println(hashedPass)
	if hashedPass != user.Pass {
		return nil, errors.New("password incorrect")
	}
	currentTime := time.Now()
	claims := dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(720 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
		Username:    user.Username,
		UserID:      user.ID,
		UserUUID:    user.UUID.String(),
		Permissions: user.Permission,
	}
	accessToken, err := s.jwtUtil.GenerateToken(&claims)
	if err != nil {
		log.Errorw(ctx, "error when generating token for user", "err", err)
		return nil, err
	}
	return &dto.LoginResponse{
		Data: &dto.Token{
			AccessToken: accessToken,
		},
	}, nil
}

func (s *authServiceImpl) authCodeSymKeyGenerate(tenantId string) (string, error) {
	if utils.IsEmpty(tenantId) {
		return "", errors.New("tenantId is empty")
	}
	rotationStep := int(tenantId[len(tenantId)-2])

	keyFull := hex.EncodeToString([]byte(utils.LeftRatation(tenantId, rotationStep)))
	if len(keyFull) < 64 {
		for i := 0; i < 64-len(keyFull); i++ {
			keyFull += "0"
		}
	}
	fmt.Println("chay duoc khong zay6")
	return keyFull[:64], nil
}
