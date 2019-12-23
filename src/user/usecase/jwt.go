package usecase

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/models"
)

func (*userUsecase) generateTokenPair(id int64) (*models.TokenPair, error) {
	jwtkey := []byte(conf.C.Auth.Jwtkey)
	expSec := conf.C.Auth.TokenExpSec
	rtExpSec := conf.C.Auth.RtExpSec
	strID := strconv.FormatInt(id, 10)

	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := jwt.StandardClaims{
		Issuer:    "idToken",
		Subject:   strID,
		ExpiresAt: expTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tString, err := token.SignedString(jwtkey)
	if err != nil {
		return nil, err
	}

	rtExpTime := time.Now().Add(time.Duration(rtExpSec) * time.Second)

	rtClaims := jwt.StandardClaims{
		Issuer:    "refreshToken",
		Subject:   strID,
		ExpiresAt: rtExpTime.Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtString, err := rt.SignedString(jwtkey)
	if err != nil {
		return nil, err
	}
	return &models.TokenPair{
		IDToken:      tString,
		RefreshToken: rtString,
	}, nil
}

func (*userUsecase) validateToken(t string) (int64, error) {
	jwtkey := []byte(conf.C.Auth.Jwtkey)
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
