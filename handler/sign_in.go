package handler

import (
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

func signIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}
		inputpw := user.Password

		db := db.Connect()
		result := db.Find(user, "email=?", user.Email)

		// 존재하지않는 아이디일 경우
		if result.RowsAffected == 0 {
			return echo.ErrBadRequest
		}

		res := helper.CheckPasswordHash(*user.Password, *inputpw)

		// 비밀번호 검증에 실패한 경우
		if !res {
			return echo.ErrUnauthorized
		}
		// 토큰 발행
		accessToken, err := helper.CreateJWT(user.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		expiredAt := time.Now().AddDate(0, 3, 0)
		user.AccessToken = &accessToken
		user.ExpiredAt = &expiredAt
		db.Save(&user)

		if err := c.JSON(http.StatusOK, map[string]string{
			"message":      "Login Success",
			"access_token": accessToken,
		}); err != nil {
			return errors.Wrap(err, "signIn")
		}
		return nil
	}
}
