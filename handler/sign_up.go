package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

// @Summary 회원가입 API
// @Description email, password를 받아 가입합니다.
// @Accept json
// @Produce json
// @Param email body string true "사용자 이메일"
// @Param password body string true "비밀번호"
// @Success 200 {object} message
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /sign-up [post]
func signUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, message{
				"bad request",
			})
		}
		db := db.Connect()
		result := db.Find(&user, "email=?", user.Email)

		// 이미 이메일이 존재할 경우의 처리
		if result.RowsAffected != 0 {
			return c.JSON(http.StatusBadRequest, message{
				"existing email",
			})
		}

		// 비밀번호를 bycrypt 라이브러리로 해싱 처리
		hashpw, err := helper.HashPassword(*user.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, message{
				err.Error(),
			})
		}
		user.Password = &hashpw

		// 위의 두단계에서 err가 nil일 경우 DB에 유저를 생성
		if err := db.Create(&user); err.Error != nil {
			return c.JSON(http.StatusInternalServerError, message{
				"Failed SignUp",
			})
		}

		// 모든 처리가 끝난 후 200, Success 메시지를 반환
		if err := c.JSON(http.StatusOK, message{
			"Success",
		}); err != nil {
			return errors.Wrap(err, "signup")
		}
		return nil
	}
}
