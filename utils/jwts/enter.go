package jwts

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	NickName string `json:"nick_name"`
	//UserName string     `json:"user_name"`
	Role   int    `json:"role"`
	Avatar string `json:"avatar"`
}

var JwtSecretKey []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
