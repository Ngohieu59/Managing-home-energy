package dto

import "github.com/golang-jwt/jwt/v5"

type LoginResponse struct {
	Data *Token `json:"data"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserID      uint   `json:"user_id"`
	Username    string `json:"user_username"`
	UserUUID    string `json:"user_uuid"`
	Permissions string `json:"user_permission"`
}

type PasswordLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
