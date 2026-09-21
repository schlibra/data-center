package utils

import (
	"data-center/internal/models"
	"errors"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtUserGenerate(userId int, username string, admin int, enable int, tokenId string) (string, error) {
	now := time.Now()
	cfg, err := LoadConfig()
	if err != nil {
		return "", err
	}
	sk := []byte(cfg.Jwt.Key)
	claims := models.JwtToken{
		UserID:   userId,
		Username: username,
		Admin:    admin,
		Enable:   enable,
		TokenID:  tokenId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "frp-auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(sk)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
func JwtUserCheck(tokenStr string) (*models.JwtToken, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	sk := []byte(cfg.Jwt.Key)
	token, err := jwt.ParseWithClaims(tokenStr, &models.JwtToken{}, func(token *jwt.Token) (any, error) {
		// 必须检查签名算法，防御 "alg: none" 或算法混淆攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return sk, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token 已过期")
		}
		return nil, fmt.Errorf("token 无效: %w", err)
	}

	// 提取并校验有效载荷
	if claims, ok := token.Claims.(*models.JwtToken); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无法解析 token claims")
}
