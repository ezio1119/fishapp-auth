package interactor

import (
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-auth/conf"
	"github.com/ezio1119/fishapp-auth/domain"
	"github.com/google/uuid"
)

type TokenInteractor struct{}

type UTokenInteractor interface {
	GenerateTokenPair(id int64) (*domain.TokenPair, error)
}

func (*TokenInteractor) GenerateTokenPair(id int64) (*domain.TokenPair, error) {
	expSec := conf.C.Auth.IDTokenExpSec
	rtExpSec := conf.C.Auth.RtExpSec
	strID := strconv.FormatInt(id, 10)

	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Subject:   strID,
		ExpiresAt: expTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tString, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	rtExpTime := time.Now().Add(time.Duration(rtExpSec) * time.Second)

	rtClaims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Subject:   strID,
		ExpiresAt: rtExpTime.Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodES256, rtClaims)
	rtString, err := rt.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return &domain.TokenPair{
		IDToken:      tString,
		RefreshToken: rtString,
	}, nil
}

var privateKey *ecdsa.PrivateKey

func init() {
	var err error
	data := []byte(conf.C.Auth.PvtJwtkey)
	if conf.C.Sv.Debug {
		// 開発環境はpemから読み込む
		data, err = ioutil.ReadFile("./dev_pvt_jwtkey.pem")
		if err != nil {
			log.Fatal(err)
		}
	}
	privateKey, err = jwt.ParseECPrivateKeyFromPEM(data)
	if err != nil {
		log.Fatal(err)
	}
}
