package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func CreateJWT(Email string) (string, error) {
	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", errors.Wrap(err, "aToken.SignedString")
	}
	return tk, nil
}

func ValidateJWT() (bool, error) {
	return false, nil
}
