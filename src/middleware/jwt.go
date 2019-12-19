package middleware

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-user/conf"
)

func (m *GoMiddleware) parseToken(t string) (int64, error) {
	jwtkey := conf.C.Auth.Jwtkey
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return 0, err
	}
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
