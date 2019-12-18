package usecase

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-user/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func (*userUsecase) GenerateTokenPair(id int64) (*models.TokenPair, error) {
	strID := strconv.FormatInt(id, 10)
	expSec, _ := strconv.ParseInt(os.Getenv("EXP_SEC"), 10, 64)
	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := jwt.StandardClaims{
		Issuer:    "idToken",
		Subject:   strID,
		ExpiresAt: expTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	rtExpSec, _ := strconv.ParseInt(os.Getenv("RT_EXP_SEC"), 10, 64)
	rtExpTime := time.Now().Add(time.Duration(rtExpSec) * time.Second)

	rtClaims := jwt.StandardClaims{
		Issuer:    "refreshToken",
		Subject:   strID,
		ExpiresAt: rtExpTime.Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtString, err := rt.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	return &models.TokenPair{
		IDToken:      tString,
		RefreshToken: rtString,
	}, nil
}
