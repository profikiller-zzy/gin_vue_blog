package jwts

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// GenerateToken 根据用户的用户名产生token
func GenerateToken(payLoad JwtPayLoad) (string, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	// Token的有效时间
	expireTime := time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.ExpireTime))
	Claim := CustomClaims{
		JwtPayLoad: payLoad,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// Issuer表示Token的签发者
			Issuer: "profikiller",
		},
	}
	// NewWithClaims根据Claims结构体创建Token示例
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	// SignedString 方法根据传入的空接口类型参数key，返回完整的签名令牌。
	reqToken, err := reqClaim.SignedString(JwtSecretKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("jwt token生成失败，错误信息：%s", err.Error()))
	}
	return reqToken, nil
}

// VerifyToken 解析和验证token
func VerifyToken(tokenString string) (*CustomClaims, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		global.Log.Error(fmt.Sprintf("Verify token error : %s", err.Error()))
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("Invalid token string")
}
