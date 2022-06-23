package helper

import (
	"os"
	"strings"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func CreateJWT(Email string) (string, error) {
	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Hour * 720 * 3).Unix()

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", errors.Wrap(err, "aToken.SignedString")
	}
	return tk, nil
}

func ValidateJWT(c echo.Context) (*models.User, error) {
	header := c.Request().Header
	authv := header.Get("Authorization")

	if authv == "" {
		return nil, errors.New("no authorization")
	}
	// Get bearer token
	if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
		return nil, errors.New("invalid bearer token")
	}

	values := strings.Split(authv, " ")
	if len(values) < 2 {
		return nil, errors.New("no bearer token")
	}

	token := values[1]
	user := new(models.User)
	db := db.Connect()
	result := db.Find(user, "access_token=?", token)
	// 존재하지않는 아이디일 경우
	if result.RowsAffected == 0 {
		return nil, echo.ErrBadRequest
	}
	if user.ExpiredAt.Before(time.Now()) {
		return nil, echo.ErrUnauthorized
	}

	return user, nil
}
