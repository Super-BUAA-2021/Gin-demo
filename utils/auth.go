package utils

import (
	"errors"
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

// GenerateToken 生成一个token
func GenerateToken(id uint64) (signedToken string) {
	expiresSeconds, err := strconv.ParseInt(global.VP.GetString("jwt.timeout"), 10, 64)
	if err != nil {
		panic("GenerateToken: Parse Timeout error")
	}
	claims := jwt.StandardClaims{
		Issuer:    "demo-server",
		ExpiresAt: expiresSeconds + time.Now().Unix(),
		Audience:  strconv.FormatUint(id, 10),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := global.VP.GetString("jwt.secret")
	signedToken, err = token.SignedString([]byte(secret))
	if err != nil {
		panic("GenerateToken: sign token error")
	}
	return
}

// ParseToken 验证token的正确性，正确则返回id
func ParseToken(signedToken string) (id uint64, err error) {
	secret := global.VP.GetString("jwt.secret")
	token, err := jwt.Parse(
		signedToken,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil || !token.Valid {
		err = errors.New("token isn't valid")
		return
	}
	id, err = strconv.ParseUint(token.Claims.(jwt.MapClaims)["aud"].(string), 10, 64)
	if err != nil {
		err = errors.New("token isn't valid")
	}
	return
}
