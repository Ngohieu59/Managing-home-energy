package utils

import (
	"Managing-home-energy/conf"
	"crypto/rsa"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

const RS256 = "RS256"

type JWTUtil interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ParseToken(token string, claims jwt.Claims) error
}

type jwtUtil struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTUtil(di *do.Injector) (JWTUtil, error) { // factory function
	cf := do.MustInvoke[*conf.Config](di)
	jwtPubkey, err := os.ReadFile(cf.JWT.PublicKeyFilePath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPubkey)
	if err != nil {
		panic(err)
	}
	jwtPrivateKey, err := os.ReadFile(cf.JWT.PrivateKeyFilePath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKey)
	if err != nil {
		return nil, err
	}
	jwtUtl := &jwtUtil{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return jwtUtl, nil
}

func (h *jwtUtil) GenerateToken(claims jwt.Claims) (string, error) { // pointer jwtUtil implement interface JWTUtil
	if claims == nil {
		return "", errors.New("claims must not be nil")
	}
	tkn := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claims)
	str, err := tkn.SignedString(h.privateKey)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (h *jwtUtil) ParseToken(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}

// Use pointer help can change data of struct, avoid copy struct.
