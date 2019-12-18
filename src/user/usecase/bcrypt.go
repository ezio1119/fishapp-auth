package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

func (*userUsecase) genEncryptedPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (*userUsecase) compareHashAndPass(encryptedPass string, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(pass)); err != nil {
		return err
	}
	return nil
}
