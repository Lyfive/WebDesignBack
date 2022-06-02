/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package util

import (
	"github.com/golang-jwt/jwt"
	"time"
	"webDesign/models"
	setting "webDesign/pkg"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "LyFive",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


func CheckToken(token string) bool {
	if token == "" {
		return false
	}
	claims, err := ParseToken(token)
	if err != nil {
		return false
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}
	return true
}

func GetUserLevel(token string) int {
	claims, err := ParseToken(token)
	if err != nil {
		return 4
	}
	return models.GetUser(claims.Username).Level
}
