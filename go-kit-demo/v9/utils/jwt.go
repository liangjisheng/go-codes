package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("jwtSecret_v9")

// JwtContextKey ...
const JwtContextKey = "jwt_context_key"

// Token ...
type Token struct {
	Name string
	DcID int
	jwt.StandardClaims
}

// CreateJwtToken ...
func CreateJwtToken(name string, dcID int) (string, error) {
	var token Token
	token.StandardClaims = jwt.StandardClaims{
		Audience:  "",                                      // 受众群体
		ExpiresAt: time.Now().Add(30 * time.Second).Unix(), // 到期时间
		Id:        "",                                      // 编号
		IssuedAt:  time.Now().Unix(),                       // 签发时间
		Issuer:    "go-kit-v9",                             // 签发人
		NotBefore: time.Now().Unix(),                       // 生效时间
		Subject:   "login",                                 // 主题
	}
	token.Name = name
	token.DcID = dcID
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken ...
func ParseToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || jwtToken == nil {
		return nil, err
	}
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		return claim, nil
	}
	return nil, nil
}
