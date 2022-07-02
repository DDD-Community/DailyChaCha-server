package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type email struct {
	Email string `json:"email"`
}

type message struct {
	Message string `json:"message"`
}

// @Summary auth 토큰으로 테스트해볼 API입니다.
// @Description access token을 확인하여 해당 토큰 유저의 이메일을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} email
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /getlist [get]
func healthCheck(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		if err := c.JSON(http.StatusOK, email{
			Email: chaUser.Email,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
