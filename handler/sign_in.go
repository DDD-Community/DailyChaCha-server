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

// @Summary 로그인 API
// @Description email, password를 받아 access token을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body models.User true "The input todo struct"
// @Success 200 {object} models.Auth
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /sign-in [post]
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

		if err := c.JSON(http.StatusOK, models.Auth{
			AccessToken: accessToken,
			ExpiredAt:   user.ExpiredAt.Format("2006-01-02"),
		}); err != nil {
			return errors.Wrap(err, "signIn")
		}
		return nil
	}
}
