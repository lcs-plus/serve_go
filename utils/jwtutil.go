package utils

import (
	"0121_1/global"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type MyCustomClaims struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte(global.App.Config.Jwt.Secretary)

func GetCacheToken(id int, role string) (string, error) {
	c := context.Background()

	key := "userToken_" + strconv.Itoa(id)
	token, err := global.App.Redis.Get(c, key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	if token == "" {
		token, err = GetToken(id, role)
		global.App.Redis.Set(c, key, token, time.Duration(global.App.Config.Jwt.ExpiresAt)*time.Second)
	}

	return token, nil
}

func GetToken(id int, role string) (string, error) {
	claims := MyCustomClaims{
		id,
		role,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			//global.App.Config.Jwt.ExpiresAt
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.ExpiresAt) * time.Second)), //过期时间，是个可变参数
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    global.App.Config.Jwt.Issuer,
			Subject:   global.App.Config.Jwt.Subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 定义解析函数，用于验证签名密钥
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否为预期的 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// 返回签名密钥
		return mySigningKey, nil
	})

	if err != nil {
		// 记录错误日志
		global.App.Log.Error(fmt.Sprintf("Failed to parse JWT: %v", err))
		return nil, err
	}

	// 检查解析后的 token 是否有效
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
