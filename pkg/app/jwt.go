package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/tool"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, time.Time, error) {
	currentTime := time.Now()
	expireTime := currentTime.Add(global.JwtSetting.Expire)

	claims := Claims{
		AppKey:    tool.EncodeMd5(appKey),
		AppSecret: tool.EncodeMd5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JwtSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(GetJwtSecret())

	return token, expireTime, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
