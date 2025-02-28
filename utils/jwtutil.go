package utils

import (
	"0121_1/global"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetToken(id int, username string) (string, error) {
	claims := MyCustomClaims{
		id,
		username, //根据用户名来动态生成
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time

			//global.App.Config.Jwt.ExpiresAt
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.ExpiresAt))), //过期时间，是个可变参数
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    global.App.Config.Jwt.Issuer,
			Subject:   global.App.Config.Jwt.Subject,
		},
	}

	var mySigningKey = []byte(global.App.Config.Jwt.Secretary)

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
		return []byte(global.App.Config.Jwt.Secretary), nil
	})

	if err != nil {
		// 记录错误日志
		log.Printf("Failed to parse JWT: %v", err)
		return nil, err
	}

	// 检查解析后的 token 是否有效
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
