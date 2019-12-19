package usecase

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/models"
)

func (*userUsecase) GenerateTokenPair(id int64) (*models.TokenPair, error) {
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
